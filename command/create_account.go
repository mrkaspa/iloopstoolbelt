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

func CreateAccountCMD(c *cli.Context) {
	userLogin := models.UserLogin{
		Email:    c.String("email"),
		Password: c.String("password"),
	}
	if err := CreateAccount(&userLogin); err == nil {
		fmt.Println("Your user account has been created, try logging in")
	} else {
		fmt.Println(err)
	}
}

func CreateAccount(userLogin *models.UserLogin) error {
	if valid, errMap := models.ValidStruct(userLogin); valid {
		userJSON, _ := json.Marshal(userLogin)
		resp, _ := client.CallRequest("POST", "/users", bytes.NewReader(userJSON))
		if resp.StatusCode == http.StatusOK {
			return nil
		} else {
			return errors.New("There was an error creating that account, please try again")
		}
	} else {
		fmt.Println(errMap)
		return errMap
	}
}
