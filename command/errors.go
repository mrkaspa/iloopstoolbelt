package command

import "errors"

var (
	ErrAccountNotCreated          = errors.New("There was an error creating that account, please try again")
	ErrProjectNotFound            = errors.New("Could not find the project")
	ErrProjectOrUserNotFound      = errors.New("Could not find the project or the user")
	ErrProjectNotDeleted          = errors.New("Could not delete the project")
	ErrProjectNotCreated          = errors.New("There was an error creating that project, please try again")
	ErrProjectNotAccess           = errors.New("You don't have permission to execute this on the project")
	ErrSSHFileNotFound            = errors.New("SSH File not found")
	ErrSSHNotCreated              = errors.New("Could not add the ssh, please try uploading another ssh")
	ErrWithCredentials            = errors.New("There was an error with the credentials, please try again")
	ErrProjectUserNotAdded        = errors.New("The user could not be added to the project")
	ErrProjectUserNotRemoved      = errors.New("The user could not be removed from the project")
	ErrProjectUserNotDelegated    = errors.New("The user could not be assigned as the admin of the project")
	ErrProjectNotLeft             = errors.New("The user could not be leave the project")
	ErrNoActiveSession            = errors.New("There is not an active session, please login")
	ErrProjectScriptNotFound      = errors.New("The script file has not been found")
	ErrProjectNameRequired        = errors.New("The project name is required")
	ErrProjectDirNotFound         = errors.New("The directory does not exist")
	ErrProjectAlreadyInit         = errors.New("The project is already initialized")
	ErrProjectAlreadyRemoteILoops = errors.New("The project already has an iloops remote branch")
	ErrArgNotFound                = errors.New("The argument wasn't found")
)
