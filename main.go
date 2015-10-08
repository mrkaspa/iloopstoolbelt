package main

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codegangsta/cli"
)

var (
	URL     = "http://api.infiniteloops.co:8080"
	VERSION = "1.1.5"
)

func main() {
	command.Init(URL)
	app := cli.NewApp()
	app.Name = "toolbelt"
	app.Usage = "ILoops command client will allow you to deploy projects on the cloud"
	app.Version = VERSION
	setCommands(app)
	app.Run(os.Args)
}

func setCommands(app *cli.App) {
	app.Commands = []cli.Command{
		command.CreateAccountCMD,
		command.LoginCMD,
		command.LogoutCMD,
		command.SSHAddCMD,
		command.SSHListCMD,
		command.SSHDeleteCMD,
		command.ProjectInitCMD,
		command.ProjectListCMD,
		command.ProjectDeleteCMD,
		command.ProjectLeaveCMD,
		command.ProjectUserAddCMD,
		command.ProjectUserRemoveCMD,
		command.ProjectUserDelegateCMD,
	}
}
