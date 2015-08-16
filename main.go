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
	app.Name = "iloops"
	app.Usage = "ILoops command client will allow you to deploy projects on the cloud"
	app.Version = "1.0.0"

	setFlags(app)
	setCommands(app)

	app.Run(os.Args)
}

func setFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "email, e",
			Usage: "user email",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "user password",
		},
	}
}

func setCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:   "create",
			Usage:  "creates a new account",
			Action: command.CreateAccountCMD,
		},
		{
			Name:   "login",
			Usage:  "login with an account",
			Action: command.Login,
		},
		{
			Name:   "logout",
			Usage:  "logout from an account",
			Action: command.Logout,
		},
	}
}
