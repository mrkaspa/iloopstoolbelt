package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"bitbucket.org/kiloops/api/endpoint"
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/api/utils"
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
	var jError endpoint.JSONError
	return client.CallRequest("POST", "/users/login", bytes.NewReader(userJSON)).Solve(utils.MapExec{
		http.StatusOK: utils.InfoExec{
			&user,
			func(resp *http.Response) error {
				return LoginFile(&user)
			},
		},
		http.StatusConflict: utils.InfoExec{
			&jError,
			func(resp *http.Response) error {
				return jError
			},
		},
		utils.Default: utils.InfoExec{
			nil,
			func(resp *http.Response) error {
				return ErrWithCredentials
			},
		},
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
