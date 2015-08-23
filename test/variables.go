package test

import "bitbucket.org/kiloops/api/models"

var (
	SSHPath          = "/Users/michelperez/.ssh/id_rsa.pub"
	anotherSSHPath   = "/Users/michelperez/.ssh/github_rsa.pub"
	userLogin        = models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
	anotherUserLogin = models.UserLogin{Email: "angelbotto@gmail.com", Password: "h1h1h1h1h1h1"}
	project          = models.Project{Name: "demo"}
)
