package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"bitbucket.org/kiloops/api/models"
	"github.com/codegangsta/cli"
)

//LoginCMD command
var LoginCMD = cli.Command{
	Name:   "login",
	Usage:  "login with credentials",
	Flags:  []cli.Flag{emailFlag, passwordFlag},
	Action: loginImpl,
}

func loginImpl(c *cli.Context) {
	userLogin := models.UserLogin{
		Email:    c.String("email"),
		Password: c.String("password"),
	}
	for userLogin.Password == "" {
		fmt.Println("Enter password: ")
		var in string
		fmt.Scanln(&in)
		userLogin.Password = in
	}
	if err := Login(&userLogin); err == nil {
		fmt.Println("Welcome!")
	} else {
		PrintError(err)
	}
}

//Login an user
func Login(userLogin *models.UserLogin) error {
	Logout()
	if valid, errMap := models.ValidStruct(userLogin); !valid {
		return errMap
	}
	userJSON, _ := json.Marshal(userLogin)
	var user models.UserLogged
	return client.CallRequest("POST", "/users/login", bytes.NewReader(userJSON)).WithResponseJSON(&user, func(resp *http.Response) error {
		switch resp.StatusCode {
		case http.StatusOK:
			return LoginFile(&user)
		case http.StatusBadRequest:
			return ErrWithCredentials
		default:
			return nil
		}
	})
}

//LoginFile configuration file
func LoginFile(user *models.UserLogged) error {
	if err := os.Mkdir(InfiniteFolder(), os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	if _, err := os.Create(InfiniteConfigFile()); err != nil {
		return err
	}
	authJSON, _ := json.Marshal(user)
	return ioutil.WriteFile(InfiniteConfigFile(), []byte(authJSON), os.ModePerm)
}
