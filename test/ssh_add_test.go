package test

import (
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSHAdd", func() {

	var (
		SSHPath        = "/Users/michelperez/.ssh/id_rsa.pub"
		anotherSSHPath = "/Users/michelperez/.ssh/github_rsa.pub"
		userLogin      = models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
	)

	BeforeEach(func() {
		cleanDB()
		command.CreateAccount(&userLogin, SSHPath)
		command.Login(&userLogin)
	})

	It("uploads a new SSH", func() {
		err := command.SSHAdd("Another ssh", anotherSSHPath)
		Expect(err).To(BeNil())
	})

})
