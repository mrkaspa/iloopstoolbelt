package command

import (
	"fmt"
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
	if err := Logout(); err == nil {
		fmt.Println("Good bye!")
	} else {
		PrintError(err)
	}
}

//Logout an account
func Logout() error {
	return os.RemoveAll(InfiniteFolder())
}
