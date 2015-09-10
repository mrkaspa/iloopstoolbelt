package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"bitbucket.org/kiloops/api/ierrors"
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/api/utils"

	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/gosimple/slug"
	"github.com/mrkaspa/go-helpers"
)

//ProjectCreateCMD command
var ProjectCreateCMD = cli.Command{
	Name:   "project:create",
	Usage:  "creates a new project with the given name",
	Action: projectCreateImpl,
}

func projectCreateImpl(c *cli.Context) {
	project := models.Project{Name: c.Args()[0]}
	if err := ProjectCreate(&project); err == nil {
		fmt.Println("Start to hack :)")
	} else {
		PrintError(err)
	}
}

//ProjectCreate new
func ProjectCreate(project *models.Project) error {
	name := slug.Make(project.Name)
	if helpers.FileExists(name) {
		return ErrProjectDirectoryAlreadyExists
	}
	return withUserSession(func(user *models.UserLogged) error {
		if valid, errMap := models.ValidStruct(project); !valid {
			return errMap
		}
		projectJSON, _ := json.Marshal(project)
		var appError ierrors.AppError
		return client.CallRequestWithHeaders("POST", "/projects", bytes.NewReader(projectJSON), authHeaders(user)).Solve(utils.MapExec{
			http.StatusOK: utils.InfoExec{
				&project,
				func(resp *http.Response) error {
					return initProject(project, name)
				},
			},
			http.StatusConflict: utils.InfoExec{
				&appError,
				func(resp *http.Response) error {
					return appError
				},
			},
			utils.Default: utils.InfoExec{
				nil,
				func(resp *http.Response) error {
					return ErrProjectNotCreated
				},
			},
		})
	})
}

func initProject(project *models.Project, name string) error {
	slug := project.Slug
	git := project.URLRepo
	// Clone project
	fmt.Println("Cloning basic project")
	err := sh.Command("git", "clone", DefaultURLProject, name).Run()
	if err != nil {
		return err
	}
	iloopProject, _ := ioutil.ReadFile(IDLoopProjectFileConfig(name))
	iloopPackage, _ := ioutil.ReadFile(IDLoopProjectPackage(name))
	overrideFile(iloopProject, IDLoopProjectFileConfig(name), name, slug)
	overrideFile(iloopPackage, IDLoopProjectPackage(name), name, name)
	return sh.NewSession().SetDir(name).Command("git", "remote", "set-url", "origin", git).Run()

}

func overrideFile(file []byte, path string, name string, id string) {
	template := string(file)
	iloopProjectContent := fmt.Sprintf(template, name, id)
	ioutil.WriteFile(path, []byte(iloopProjectContent), os.ModePerm)
}
