package command

import "github.com/codegangsta/cli"

func LoginCMD() cli.Command {
	return cli.Command{
		Name:  "login",
		Usage: "login with credentials",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email, e",
				Usage: "user email",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "user password",
			},
		},
		Action: func(c *cli.Context) {

		},
	}
}
