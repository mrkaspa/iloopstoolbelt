package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

//ProjectGetCMD command
var ProjectGetCMD = cli.Command{
	Name:   "project:get",
	Usage:  "downloads a project",
	Action: projectGetImpl,
}

func projectGetImpl(c *cli.Context) {
	if err := ProjectGet(); err == nil {
		fmt.Println("Start to hack :)")
	} else {
		PrintError(err)
	}
}

//ProjectGet a previous one
func ProjectGet() error {
	return nil
}
