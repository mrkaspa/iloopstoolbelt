package main

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codegangsta/cli"
)

var (
	URL = "http://localhost:8080"
)

func main() {
	command.Init(URL)
	app := cli.NewApp()
	app.Name = "toolbelt"
	app.Usage = "ILoops command client will allow you to deploy projects on the cloud"
	app.Version = "1.0.0"
	setCommands(app)
	app.Run(os.Args)
}

func setCommands(app *cli.App) {
	app.Commands = []cli.Command{
		command.CreateAccountCMD,
		command.LoginCMD,
		command.LogoutCMD,
		command.SSHAddCMD,
		command.ProjectCreateCMD,
		command.ProjectListCMD,
		command.ProjectDeleteCMD,
	}
}
