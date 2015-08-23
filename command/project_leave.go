package command

import (
	"bytes"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"
	"github.com/codegangsta/cli"
)

//ProjectUserAddCMD command
var ProjectLeaveCMD = cli.Command{
	Name:   "project:leave",
	Usage:  "leaves a project",
	Action: projectUserAddImpl,
}

func projectLeaveImpl(c *cli.Context) {
	slug := c.Args()[0]
	if err := ProjectLeave(slug); err == nil {
		fmt.Println("You have left the project")
	} else {
		PrintError(err)
	}
}

//ProjectUserAdd an account
func ProjectLeave(slug string) error {
	return withUserSession(func(user *models.UserLogged) error {
		resp, _ := client.CallRequestWithHeaders("PUT", "/projects/"+slug+"/leave", bytes.NewReader(emptyJSON), authHeaders(user))
		switch resp.StatusCode {
		case http.StatusOK:
			return nil
		case http.StatusBadRequest:
			return ErrProjectNotLeft
		case http.StatusNotFound:
			return ErrProjectNotFound
		case http.StatusForbidden:
			return ErrProjectNotAccess
		}
		return nil
	})
}
