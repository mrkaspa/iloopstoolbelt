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

//ChangePasswordCMD command
var ChangePasswordCMD = cli.Command{
	Name:   "password:change",
	Usage:  "sets a new password",
	Action: changePasswordImpl,
}

func changePasswordImpl(c *cli.Context) {
	changePassword := models.ChangePassword{}
	if changePassword.Token == "" {
		changePassword.Token = readLine("Enter the token:")
	}
	changePassword.Password = readPassword("Enter the new password:")
	secondPassword := readPassword("Enter the new password again:")
	if changePassword.Password != secondPassword {
		PrintError(ErrPasswordsUnMatch)
		return
	}
	if err := ChangePassword(&changePassword); err == nil {
		fmt.Println("An email will be sent to your email")
	} else {
		PrintError(err)
	}
}

//ChangePassword of an user
func ChangePassword(changePassword *models.ChangePassword) error {
	if valid, errMap := models.ValidStruct(changePassword); !valid {
		return errMap
	}
	changePasswordJSON, _ := json.Marshal(changePassword)
	var appError ierrors.AppError
	return client.CallRequest("POST", "/users/change_password", bytes.NewReader(changePasswordJSON)).Solve(utils.MapExec{
		http.StatusOK: utils.InfoExec{
			nil,
			func(resp *http.Response) error {
				return nil
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
