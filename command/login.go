package command

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"bitbucket.org/kiloops/api/models"
	"github.com/codegangsta/cli"
)

//LoginCMD command
var LoginCMD = cli.Command{
	Name:  "login",
	Usage: "login with credentials",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "email, e",
			Usage: "user email",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "user password",
		},
	},
	Action: loginImpl,
}

func loginImpl(c *cli.Context) {
	userLogin := models.UserLogin{
		Email:    c.String("email"),
		Password: c.String("password"),
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
	if valid, errMap := models.ValidStruct(userLogin); valid {
		userJSON, _ := json.Marshal(userLogin)
		resp, _ := client.CallRequest("POST", "/users/login", bytes.NewReader(userJSON))
		defer resp.Body.Close()
		switch resp.StatusCode {
		case http.StatusOK:
			var user models.UserLogged
			GetBodyJSON(resp, &user)
			return LoginFile(&user)
		case http.StatusBadRequest:
			return errors.New("There was an error with the credentials, please try again")
		}
	} else {
		return errMap
	}
	return nil
}

//LoginFile configuration file
func LoginFile(user *models.UserLogged) error {
	if err := os.Mkdir(InfiniteFolder(), os.ModePerm); err == nil || os.IsExist(err) {
		if _, err := os.Create(InfiniteConfigFile()); err == nil {
			authJSON, _ := json.Marshal(user)
			return ioutil.WriteFile(InfiniteConfigFile(), []byte(authJSON), os.ModePerm)
		} else {
			return err
		}
	} else {
		return err
	}
}
