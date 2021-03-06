package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrkaspa/iloopsapi/ierrors"
	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopsapi/utils"
	"github.com/codegangsta/cli"
)

//LoginCMD command
var LoginCMD = cli.Command{
	Name:   "login",
	Usage:  "login with credentials",
	Flags:  []cli.Flag{emailFlag},
	Action: loginImpl,
}

func loginImpl(c *cli.Context) {
	userLogin := models.UserLogin{
		Email: c.String("email"),
	}
	if userLogin.Email == "" {
		userLogin.Email = readLine("Enter your email:")
	}
	userLogin.Password = readPassword("Enter your password:")
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
	var appError ierrors.AppError
	return client.CallRequest("POST", "/users/login", bytes.NewReader(userJSON)).Solve(utils.MapExec{
		http.StatusOK: utils.InfoExec{
			&user,
			func(resp *http.Response) error {
				return loginFile(&user)
			},
		},
		http.StatusConflict: utils.InfoExec{
			&appError,
			func(resp *http.Response) error {
				return appError
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
