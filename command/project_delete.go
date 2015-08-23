package command

import (
	"bytes"
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
	if err := ProjectDelete(c.Args()[0]); err == nil {
		fmt.Println("The project has been deleted")
	} else {
		PrintError(err)
	}
}

//ProjectDelete an account
func ProjectDelete(slug string) error {
	return withUserSession(func(user *models.UserLogged) error {
		resp, _ := client.CallRequestWithHeaders("DELETE", "/projects/"+slug, bytes.NewReader(emptyJSON), authHeaders(user))
		switch resp.StatusCode {
		case http.StatusOK:
			return nil
		case http.StatusBadRequest:
			return ErrProjectNotDeleted
		case http.StatusForbidden:
			return ErrProjectNotAccess
		case http.StatusNotFound:
			return ErrProjectNotFound
		}
		return nil
	})
}
