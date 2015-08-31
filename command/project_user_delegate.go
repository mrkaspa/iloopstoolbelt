package command

import (
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"
	"github.com/codegangsta/cli"
)

//ProjectUserDelegateCMD command
var ProjectUserDelegateCMD = cli.Command{
	Name:   "project:user:delegate",
	Usage:  "delegates an user as the admin of the project",
	Flags:  []cli.Flag{emailFlag},
	Action: projectUserDelegateImpl,
}

func projectUserDelegateImpl(c *cli.Context) {
	email := c.String("email")
	slug := c.Args()[0]
	if err := ProjectUserDelegate(slug, email); err == nil {
		fmt.Println("The user has been assigned as the admin of the project")
	} else {
		PrintError(err)
	}
}

//ProjectUserDelegate an account
func ProjectUserDelegate(slug, email string) error {
	return withUserSession(func(user *models.UserLogged) error {
		resp, _ := client.CallRequestNoBodytWithHeaders("PUT", "/projects/"+slug+"/delegate/"+email, authHeaders(user))
		switch resp.StatusCode {
		case http.StatusOK:
		case http.StatusBadRequest:
			return ErrProjectUserNotDelegated
		case http.StatusNotFound:
			return ErrProjectOrUserNotFound
		case http.StatusForbidden:
			return ErrProjectNotAccess
		}
		return nil
	})
}
