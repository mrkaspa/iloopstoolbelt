package command

import "github.com/codegangsta/cli"

var emailFlag = cli.StringFlag{
	Name:  "email, e",
	Usage: "user email",
}

var sshFlag = cli.StringFlag{
	Name:  "ssh, s",
	Usage: "ssh key",
}

var nameFlag = cli.StringFlag{
	Name:  "name, n",
	Usage: "name for the ssh key",
}
