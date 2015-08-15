package command

import "github.com/codegangsta/cli"

func CreateAccount(c *cli.Context) {
	println("added task: ", c.Args().First())
}
