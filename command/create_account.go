package command

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"

	"github.com/codegangsta/cli"
)

//CreateAccountCMD command
func CreateAccountCMD(c *cli.Context) {
	userLogin := models.UserLogin{
		Email:    c.String("email"),
		Password: c.String("password"),
	}
	if err := CreateAccount(&userLogin); err == nil {
		fmt.Println("Your user account has been created, try logging in")
	} else {
		PrintError(err)
	}
}

//CreateAccount new account
func CreateAccount(userLogin *models.UserLogin) error {
	if valid, errMap := models.ValidStruct(userLogin); valid {
		userJSON, _ := json.Marshal(userLogin)
		resp, _ := client.CallRequest("POST", "/users", bytes.NewReader(userJSON))
		switch resp.StatusCode {
		case http.StatusOK:
		case http.StatusBadRequest:
			return errors.New("There was an error creating that account, please try again")
		}
	} else {
		return errMap
	}
	return nil
}
