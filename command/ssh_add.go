package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrkaspa/iloopsapi/models"

	"github.com/codegangsta/cli"
	"github.com/mrkaspa/go-helpers"
)

//SSHAddCMD command
var SSHAddCMD = cli.Command{
	Name:   "ssh:add",
	Usage:  "logout from the current account",
	Flags:  []cli.Flag{nameFlag, sshFlag},
	Action: sshAddImpl,
}

func sshAddImpl(c *cli.Context) {
	SSHPath := c.String("ssh")
	name := c.String("name")
	if SSHPath == "" {
		SSHPath = readLine("Enter the ssh path:")
	}
	if name == "" {
		name = readLine("Enter the name for the key:")
	}
	if err := SSHAdd(name, SSHPath); err == nil {
		fmt.Println("The ssh key has been added")
	} else {
		PrintError(err)
	}
}

//SSHAdd new key
func SSHAdd(name, SSHPath string) error {
	if !helpers.FileExists(SSHPath) {
		return ErrSSHFileNotFound
	}
	return withUserSession(func(user *models.UserLogged) error {
		return UploadSSH(name, SSHPath, user)
	})
}

//UploadSSH key for the current user
func UploadSSH(name string, SSHPath string, user *models.UserLogged) error {
	SSHContent, _ := ioutil.ReadFile(SSHPath)
	content := string(SSHContent)
	ssh := models.SSH{Name: name, PublicKey: content}
	sshJSON, _ := json.Marshal(ssh)
	return client.CallRequestWithHeaders("POST", "/ssh", bytes.NewReader(sshJSON), authHeaders(user)).WithResponse(func(resp *http.Response) error {
		if resp.StatusCode != http.StatusOK {
			return ErrSSHNotCreated
		}
		return nil
	})
}
