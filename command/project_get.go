package command

import (
	"bytes"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"

	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/gosimple/slug"
)

//ProjectGetCMD command
var ProjectGetCMD = cli.Command{
	Name:   "project:get",
	Usage:  "downloads a project",
	Action: projectGetImpl,
}

func projectGetImpl(c *cli.Context) {
	if err := ProjectGet(c.Args()[0]); err == nil {
		fmt.Println("Start to hack :)")
	} else {
		PrintError(err)
	}
}

//ProjectGet a previous one
func ProjectGet(slug string) error {
	return withUserSession(func(user *models.UserLogged) error {
		resp, _ := client.CallRequestWithHeaders("GET", "/projects/"+slug, bytes.NewReader(emptyJSON), authHeaders(user))
		var project models.Project
		switch resp.StatusCode {
		case http.StatusOK:
			GetBodyJSON(resp, &project)
			cloneProject(&project)
		case http.StatusNotFound:
			return ErrProjectOrUserNotFound
		case http.StatusForbidden:
			return ErrProjectNotAccess
		}
		return nil
	})
}

func cloneProject(project *models.Project) {
	name := slug.Make(project.Name)
	git := project.URLRepo
	// Clone project
	fmt.Println("Cloning project")
	sh.Command("git", "clone", git, name).Run()
}
