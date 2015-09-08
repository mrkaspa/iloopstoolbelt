package command

import (
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"
	"github.com/codegangsta/cli"
)

//ProjectUserAddCMD command
var ProjectUserAddCMD = cli.Command{
	Name:   "project:user:add",
	Usage:  "adds an user by email to the project",
	Flags:  []cli.Flag{emailFlag},
	Action: projectUserAddImpl,
}

func projectUserAddImpl(c *cli.Context) {
	email := c.String("email")
	slug := c.Args()[0]
	if err := ProjectUserAdd(slug, email); err == nil {
		fmt.Println("The user has been added to the project")
	} else {
		PrintError(err)
	}
}

//ProjectUserAdd an account
func ProjectUserAdd(slug, email string) error {
	return withUserSession(func(user *models.UserLogged) error {
		return client.CallRequestNoBodytWithHeaders("PUT", "/projects/"+slug+"/add/"+email, authHeaders(user)).WithResponse(func(resp *http.Response) error {
			switch resp.StatusCode {
			case http.StatusConflict:
				return ErrProjectUserNotAdded
			case http.StatusNotFound:
				return ErrProjectOrUserNotFound
			case http.StatusForbidden:
				return ErrProjectNotAccess
			default:
				return nil
			}
		})
	})
}
