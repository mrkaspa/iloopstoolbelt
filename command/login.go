package command

import "github.com/codegangsta/cli"

func Login(c *cli.Context) {
	println("added task: ", c.Args().First())
}
