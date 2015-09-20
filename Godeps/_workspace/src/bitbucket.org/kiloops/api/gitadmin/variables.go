package gitadmin

import (
	"os"

	"bitbucket.org/kiloops/api/utils"
)

var (
	GITOLITEPATH, GITURLROOT string
)

func InitVars() {
	utils.Log.Debug("*** INIT GITADMIN VARS ***")
	GITOLITEPATH = os.Getenv("GITOLITE_PATH")
	GITURLROOT = os.Getenv("GIT_URL_ROOT")
	utils.Log.Infof("GITOLITEPATH = %s", GITOLITEPATH)
	utils.Log.Infof("GITURLROOT = %s", GITURLROOT)
}
