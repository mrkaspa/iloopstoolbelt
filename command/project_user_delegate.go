package command

import (
	"fmt"
	"net/http"

	"github.com/mrkaspa/iloopsapi/models"
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
	if err := validateArgAt(c.Args(), 0); err != nil {
		PrintError(ErrProjectNameRequired)
		return
	}
	if email == "" {
		email = readLine("Enter the user email:")
	}
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
		return client.CallRequestNoBodytWithHeaders("PUT", "/projects/"+slug+"/delegate/"+email, authHeaders(user)).WithResponse(func(resp *http.Response) error {
			switch resp.StatusCode {
			case http.StatusConflict:
				return ErrProjectUserNotDelegated
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
