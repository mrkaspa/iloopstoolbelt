package command

import (
	"bytes"
	"fmt"
	"net/http"

	"bitbucket.org/kiloops/api/models"

	"github.com/codegangsta/cli"
)

//ProjectListCMD command
var ProjectListCMD = cli.Command{
	Name:   "project:list",
	Usage:  "list my projects",
	Action: projectListImpl,
}

func projectListImpl(c *cli.Context) {
	if err := ProjectList(); err == nil {
		fmt.Println("Good bye!")
	} else {
		PrintError(err)
	}
}

//ProjectList an account
func ProjectList() error {
	return withUserSession(func(user *models.UserLogged) error {
		resp, _ := client.CallRequestWithHeaders("GET", "/projects", bytes.NewReader(emptyJSON), authHeaders(user))
		var userProjects []models.UsersProjects
		switch resp.StatusCode {
		case http.StatusOK:
			GetBodyJSON(resp, &userProjects)
			printProjects(&userProjects)
		}
		return nil
	})
}

func printProjects(userProjects *[]models.UsersProjects) {
	for i, v := range *userProjects {
		fmt.Printf("%d. %s", i, v.Project.Name)
	}
}
