package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/ierrors"
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/api/utils"
	"github.com/codegangsta/cli"
)

//ForgotPasswordCMD command
var ForgotPasswordCMD = cli.Command{
	Name:   "password:forgot",
	Usage:  "request token for password change",
	Flags:  []cli.Flag{emailFlag},
	Action: forgotPasswordImpl,
}

func forgotPasswordImpl(c *cli.Context) {
	email := models.Email{
		Value: c.String("email"),
	}
	if email.Value == "" {
		email.Value = readLine("Enter your email:")
	}
	if err := ForgotPassword(&email); err == nil {
		fmt.Println("An email will be sent to your email")
	} else {
		PrintError(err)
	}
}

//ForgotPassword of an user
func ForgotPassword(email *models.Email) error {
	if valid, errMap := models.ValidStruct(email); !valid {
		return errMap
	}
	emailJSON, _ := json.Marshal(email)
	var passwordRequest models.PasswordRequest
	var appError ierrors.AppError
	return client.CallRequest("POST", "/users/forgot", bytes.NewReader(emailJSON)).Solve(utils.MapExec{
		http.StatusOK: utils.InfoExec{
			&passwordRequest,
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
