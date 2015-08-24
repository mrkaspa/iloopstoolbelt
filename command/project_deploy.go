package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

//ProjectDeployCMD command
var ProjectDeployCMD = cli.Command{
	Name:   "project:get",
	Usage:  "downloads a project",
	Action: projectDeployImpl,
}

func projectDeployImpl(c *cli.Context) {
	if err := ProjectDeploy(); err == nil {
		fmt.Println("The project has been deployed")
	} else {
		PrintError(err)
	}
}

//ProjectDeploy on ILoops servers
func ProjectDeploy() error {
	return nil
}
