package command

import "errors"

var (
	ErrAccountNotCreated = errors.New("There was an error creating that account, please try again")
	ErrProjectNotFound   = errors.New("Could not delete the project")
	ErrProjectNotDeleted = errors.New("Could not delete the project")
	ErrProjectNotCreated = errors.New("There was an error creating that project, please try again")
	ErrProjectNotAccess  = errors.New("You don't have permission to execute this on the project")
	ErrSSHFileNotFound   = errors.New("SSH File not found")
	ErrSSHNotCreated     = errors.New("Could not add the ssh, please try again")
	ErrWithCredentials   = errors.New("There was an error with the credentials, please try again")
)
