package command

import (
	"fmt"
	"net/http"

	"github.com/mrkaspa/iloopsapi/models"

	"github.com/codegangsta/cli"
)

//SSHAddCMD command
var SSHDeleteCMD = cli.Command{
	Name:   "ssh:delete",
	Usage:  "deletes the key associated with that name",
	Flags:  []cli.Flag{nameFlag},
	Action: sshDeleteImpl,
}

func sshDeleteImpl(c *cli.Context) {
	name := c.String("name")
	if name == "" {
		name = readLine("Enter the name for the key:")
	}
	if err := SSHDelete(name); err == nil {
		fmt.Println("The ssh key has been deleted")
	} else {
		PrintError(err)
	}
}

//SSHAdd new key
func SSHDelete(name string) error {
	return withUserSession(func(user *models.UserLogged) error {
		return client.CallRequestNoBodytWithHeaders("DELETE", "/ssh/"+name, authHeaders(user)).WithResponse(func(resp *http.Response) error {
			switch resp.StatusCode {
			case http.StatusConflict:
				return ErrSSHNotDeleted
			case http.StatusNotFound:
				return ErrSSHNotFound
			default:
				return nil
			}
		})
	})
}
