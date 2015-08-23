package command

import (
	"bytes"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"
	"github.com/codegangsta/cli"
)

//ProjectUserAddCMD command
var ProjectUserRemoveCMD = cli.Command{
	Name:   "project:user:remove",
	Usage:  "removes an user by email from the project",
	Flags:  []cli.Flag{emailFlag},
	Action: projectUserRemoveImpl,
}

func projectUserRemoveImpl(c *cli.Context) {
	email := c.String("email")
	slug := c.Args()[0]
	if err := ProjectUserRemove(slug, email); err == nil {
		fmt.Println("The user has been removed from the project")
	} else {
		PrintError(err)
	}
}

//ProjectUserRemove an account
func ProjectUserRemove(slug, email string) error {
	return withUserSession(func(user *models.UserLogged) error {
		resp, _ := client.CallRequestWithHeaders("DELETE", "/projects/"+slug+"/remove/"+email, bytes.NewReader(emptyJSON), authHeaders(user))
		switch resp.StatusCode {
		case http.StatusOK:
			return nil
		case http.StatusBadRequest:
			return ErrProjectUserNotRemoved
		case http.StatusNotFound:
			return ErrProjectOrUserNotFound
		case http.StatusForbidden:
			return ErrProjectNotAccess
		}
		return nil
	})
}
