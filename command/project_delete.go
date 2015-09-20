package command

import (
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"

	"github.com/codegangsta/cli"
)

//ProjectDeleteCMD command
var ProjectDeleteCMD = cli.Command{
	Name:   "project:delete",
	Usage:  "Deletes a project by name",
	Action: projectDeleteImpl,
}

func projectDeleteImpl(c *cli.Context) {
	if err := validateArgAt(c.Args(), 0); err != nil {
		PrintError(ErrProjectNameRequired)
		return
	}
	if err := ProjectDelete(c.Args()[0]); err == nil {
		fmt.Println("The project has been deleted")
	} else {
		PrintError(err)
	}
}

//ProjectDelete an account
func ProjectDelete(slug string) error {
	return withUserSession(func(user *models.UserLogged) error {
		return client.CallRequestNoBodytWithHeaders("DELETE", "/projects/"+slug, authHeaders(user)).WithResponse(func(resp *http.Response) error {
			switch resp.StatusCode {
			case http.StatusConflict:
				return ErrProjectNotDeleted
			case http.StatusForbidden:
				return ErrProjectNotAccess
			case http.StatusNotFound:
				return ErrProjectNotFound
			default:
				return nil
			}
		})
	})
}
