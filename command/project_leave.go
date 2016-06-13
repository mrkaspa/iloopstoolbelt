package command

import (
	"fmt"
	"net/http"

	"github.com/mrkaspa/iloopsapi/models"
	"github.com/codegangsta/cli"
)

//ProjectUserAddCMD command
var ProjectLeaveCMD = cli.Command{
	Name:   "project:leave",
	Usage:  "leaves a project",
	Action: projectUserAddImpl,
}

func projectLeaveImpl(c *cli.Context) {
	if err := validateArgAt(c.Args(), 0); err != nil {
		PrintError(ErrProjectNameRequired)
		return
	}
	slug := c.Args()[0]
	if err := ProjectLeave(slug); err == nil {
		fmt.Println("You have left the project")
	} else {
		PrintError(err)
	}
}

//ProjectLeave a project
func ProjectLeave(slug string) error {
	return withUserSession(func(user *models.UserLogged) error {
		return client.CallRequestNoBodytWithHeaders("PUT", "/projects/"+slug+"/leave", authHeaders(user)).WithResponse(func(resp *http.Response) error {
			switch resp.StatusCode {
			case http.StatusConflict:
				return ErrProjectNotLeft
			case http.StatusNotFound:
				return ErrProjectNotFound
			case http.StatusForbidden:
				return ErrProjectNotAccess
			default:
				return nil
			}
		})
	})
}
