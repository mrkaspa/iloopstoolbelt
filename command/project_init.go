package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/mrkaspa/iloopsapi/ierrors"
	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopsapi/utils"

	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/mrkaspa/go-helpers"
)

//ProjectCreateCMD command
var ProjectInitCMD = cli.Command{
	Name:   "project:init",
	Usage:  "inits a new project with the given name, if no name is given it will take the current directory name",
	Action: projectInitImpl,
}

var remoteName = "iloops"

func projectInitImpl(c *cli.Context) {
	var projectName string
	currentDir, _ := os.Getwd()

	//gets the name of the directory
	if len(c.Args()) == 0 {
		splits := strings.Split(currentDir, "/")
		projectName = splits[len(splits)-1]
		if helpers.FileExists(IloopProject()) {
			PrintError(ErrProjectAlreadyInit)
			return
		}
	} else {
		projectName = c.Args()[0]
		if !helpers.FileExists(projectName) {
			PrintError(ErrProjectDirNotFound)
			return
		}
		if helpers.FileExists(projectName + "/" + IloopProject()) {
			PrintError(ErrProjectAlreadyInit)
			return
		}
	}

	//gets the script name
	mainScript := readLine("enter the main script path relative to the project folder:")
	navigateToDir := false
	if !helpers.FileExists(mainScript) && !helpers.FileExists(projectName+"/"+mainScript) {
		PrintError(ErrProjectScriptNotFound)
		return
	}

	//validates if it has to navigate to dir
	if helpers.FileExists(projectName + "/" + mainScript) {
		navigateToDir = true
	}

	//validates if the iloops remote exists
	if validateRemoteExistence(projectName, navigateToDir) {
		PrintError(ErrProjectAlreadyRemoteILoops)
		return
	}

	cronFormat := readLine("enter the cron format:")
	project := models.Project{Name: projectName}
	if err := ProjectInit(&project, navigateToDir, mainScript, cronFormat); err == nil {
		fmt.Println("Start to hack :)")
	} else {
		PrintError(err)
	}
}

//ProjectInit new
func ProjectInit(project *models.Project, navigateToDir bool, mainScript, cronFormat string) error {
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
					return initProjectGit(project, navigateToDir, mainScript, cronFormat)
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

func initProjectGit(project *models.Project, navigateToDir bool, mainScript, cronFormat string) error {
	//variables
	projectName := project.Name
	slug := project.Slug
	git := project.URLRepo
	dirName := ""
	session := sh.NewSession()
	if navigateToDir {
		dirName = projectName + "/"
		session.SetDir(dirName)
	}

	//config file
	config := models.ProjectConfig{
		Name:       projectName,
		AppID:      slug,
		MainScript: mainScript,
		Loops:      models.Loops{CronFormat: cronFormat},
	}
	configJSON, _ := json.MarshalIndent(config, "", "  ")
	ioutil.WriteFile(dirName+IloopProject(), configJSON, os.ModePerm)

	//git init
	if !helpers.FileExists(dirName + ".git") {
		session.Command("git", "init").Run()
	}

	//git remote add
	return session.Command("git", "remote", "add", remoteName, git).Run()
}

func validateRemoteExistence(dirName string, navigateToDir bool) bool {
	if !helpers.FileExists(dirName + ".git") {
		return false
	}
	session := sh.NewSession()
	if navigateToDir {
		session.SetDir(dirName)
	}
	output, err := session.Command("git", "remote").Output()
	if err != nil {
		return false
	}
	lines := strings.Split(string(output), "\n")
	exist := false
	for _, line := range lines {
		if line == remoteName {
			exist = true
		}
	}
	return exist
}
