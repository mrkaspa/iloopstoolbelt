package command

import (
	"fmt"
	"net/http"

	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopsapi/utils"

	"github.com/codegangsta/cli"
)

//SSHAddCMD command
var SSHListCMD = cli.Command{
	Name:   "ssh:list",
	Usage:  "list all the user's keys",
	Action: sshListImpl,
}

func sshListImpl(c *cli.Context) {
	if err := SSHList(); err == nil {
		fmt.Println("The ssh key has been added")
	} else {
		PrintError(err)
	}
}

//SSHAdd new key
func SSHList() error {
	return withUserSession(func(user *models.UserLogged) error {
		var sshs []models.SSH
		return client.CallRequestNoBodytWithHeaders("GET", "/ssh", authHeaders(user)).Solve(utils.MapExec{
			http.StatusOK: utils.InfoExec{
				&sshs,
				func(resp *http.Response) error {
					printSSHKeys(&sshs)
					return nil
				},
			},
		})
	})
}

func printSSHKeys(sshs *[]models.SSH) {
	for i, v := range *sshs {
		fmt.Printf("%d. %s => %s...\n", i+1, v.Name, shortString(v.PublicKey, 5))
	}
}
