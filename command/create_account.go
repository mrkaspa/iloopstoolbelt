package command

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/kiloops/api/models"

	"github.com/codegangsta/cli"
	"github.com/mrkaspa/go-helpers"
)

//CreateAccountCMD command
func CreateAccountCMD() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "creates a new account",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email, e",
				Usage: "user email",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "user password",
			},
			cli.StringFlag{
				Name:  "ssh, s",
				Usage: "ssh key",
			},
		},
		Action: createAccountImpl,
	}
}

func createAccountImpl(c *cli.Context) {
	userLogin := models.UserLogin{
		Email:    c.String("email"),
		Password: c.String("password"),
	}
	SSHPath := c.String("ssh")
	if err := CreateAccount(&userLogin, SSHPath); err == nil {
		fmt.Println("Your user account has been created, try logging in")
	} else {
		PrintError(err)
	}
}

//CreateAccount new account
func CreateAccount(userLogin *models.UserLogin, SSHPath string) error {
	if valid, errMap := models.ValidStruct(userLogin); valid {
		if helpers.FileExists(SSHPath) {
			userJSON, _ := json.Marshal(userLogin)
			resp, _ := client.CallRequest("POST", "/users", bytes.NewReader(userJSON))
			defer resp.Body.Close()
			switch resp.StatusCode {
			case http.StatusOK:
				var user models.UserLogged
				GetBodyJSON(resp, &user)
				return UploadSSH(SSHPath, &user)
			case http.StatusBadRequest:
				return errors.New("There was an error creating that account, please try again")
			}
		} else {
			return errors.New("SSH File not found")
		}
	} else {
		return errMap
	}
	return nil
}

//UploadSSH key for the current user
func UploadSSH(SSHPath string, user *models.UserLogged) error {
	SSHContent, _ := ioutil.ReadFile(SSHPath)
	content := string(SSHContent)
	ssh := models.SSH{PublicKey: content}
	sshJSON, _ := json.Marshal(ssh)
	resp, _ := client.CallRequestWithHeaders("POST", "/ssh", bytes.NewReader(sshJSON), authHeaders(user))
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("Could not add the ssh, please add it")
}
