package command

import "github.com/codegangsta/cli"

func Logout(c *cli.Context) {
	println("added task: ", c.Args().First())
}
