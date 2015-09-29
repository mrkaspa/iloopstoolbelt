package test

import "bitbucket.org/kiloops/api/models"

var (
	SSHPath          = "/Users/michelperez/.ssh/demo_rsa.pub"
	anotherSSHPath   = "/Users/michelperez/.ssh/github_rsa.pub"
	TestURLProject   = "git@github.com:infiniteloopsco/SlackNodeDemo.git"
	scriptTest       = "src/app.js"
	cronTest         = "every 1m"
	userLogin        models.UserLogin
	anotherUserLogin models.UserLogin
	project          models.Project
)
