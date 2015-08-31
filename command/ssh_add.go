package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/kiloops/api/models"

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
	resp, _ := client.CallRequestWithHeaders("POST", "/ssh", bytes.NewReader(sshJSON), authHeaders(user))
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return ErrSSHNotCreated
}
