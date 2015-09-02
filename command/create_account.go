package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"

	"github.com/codegangsta/cli"
	"github.com/mrkaspa/go-helpers"
)

//CreateAccountCMD command
var CreateAccountCMD = cli.Command{
	Name:   "create",
	Usage:  "creates a new account",
	Flags:  []cli.Flag{emailFlag, passwordFlag, sshFlag},
	Action: createAccountImpl,
}

func createAccountImpl(c *cli.Context) {
	userLogin := models.UserLogin{
		Email:    c.String("email"),
		Password: c.String("password"),
	}
	SSHPath := c.String("ssh")
	for userLogin.Password == "" {
		fmt.Println("Enter password: ")
		var in string
		fmt.Scanln(&in)
		userLogin.Password = in
	}
	if err := CreateAccount(&userLogin, SSHPath); err == nil {
		fmt.Println("Your user account has been created, try logging in")
	} else {
		PrintError(err)
	}
}

//CreateAccount new account
func CreateAccount(userLogin *models.UserLogin, SSHPath string) error {
	if valid, errMap := models.ValidStruct(userLogin); !valid {
		return errMap
	}
	if !helpers.FileExists(SSHPath) {
		return ErrSSHFileNotFound
	}
	userJSON, _ := json.Marshal(userLogin)
	var user models.UserLogged
	return client.CallRequest("POST", "/users", bytes.NewReader(userJSON)).WithResponseJSON(&user, func(resp *http.Response) error {
		switch resp.StatusCode {
		case http.StatusOK:
			return UploadSSH("New Account", SSHPath, &user)
		case http.StatusBadRequest:
			return ErrAccountNotCreated
		default:
			return nil
		}
	})
}
