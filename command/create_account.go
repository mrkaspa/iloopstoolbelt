package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/user"

	"bitbucket.org/kiloops/api/ierrors"
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/api/utils"

	"github.com/codegangsta/cli"
	"github.com/mrkaspa/go-helpers"
)

//CreateAccountCMD command
var CreateAccountCMD = cli.Command{
	Name:   "account:create",
	Usage:  "creates a new account",
	Flags:  []cli.Flag{emailFlag, sshFlag},
	Action: createAccountImpl,
}

func createAccountImpl(c *cli.Context) {
	userLogin := models.UserLogin{
		Email: c.String("email"),
	}
	SSHPath := c.String("ssh")
	if userLogin.Email == "" {
		userLogin.Email = readLine("Enter your email:")
	}
	if SSHPath == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		SSHPath = usr.HomeDir + "/.ssh/id_rsa.pub"
	}
	userLogin.Password = readPassword("Enter your password:")
	secondPassword := readPassword("Enter your password again:")
	if userLogin.Password != secondPassword {
		PrintError(ErrPasswordsUnMatch)
		return
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
	var appError ierrors.AppError
	return client.CallRequest("POST", "/users", bytes.NewReader(userJSON)).Solve(utils.MapExec{
		http.StatusOK: utils.InfoExec{
			&user,
			func(resp *http.Response) error {
				err := UploadSSH("New Account", SSHPath, &user)
				return err
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
				return ErrAccountNotCreated
			},
		},
	})
}
