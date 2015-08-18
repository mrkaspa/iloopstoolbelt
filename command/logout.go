package command

import (
	"os"

	"github.com/codegangsta/cli"
)

//LogoutCMD command
var LogoutCMD = cli.Command{
	Name:   "logout",
	Usage:  "logout from the current account",
	Action: logoutImpl,
}

func logoutImpl(c *cli.Context) {
	Logout()
}

//Logout an account
func Logout() error {
	return os.RemoveAll(InfiniteFolder())
}
