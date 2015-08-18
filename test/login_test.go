package test

import (
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login", func() {

	var (
		SSHPath   = "/Users/michelperez/.ssh/id_rsa.pub"
		userLogin = models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
	)

	BeforeEach(func() {
		cleanDB()
		command.CreateAccount(&userLogin, SSHPath)
	})

	It("logins with a created account", func() {
		err := command.Login(&userLogin)
		Expect(err).To(BeNil())
		Expect(helpers.FileExists(command.InfiniteConfigFile())).To(BeTrue())
	})

	It("logins with a invalid credentials", func() {
		userLogin.Email = "michel.ing"
		err := command.Login(&userLogin)
		Expect(err).NotTo(BeNil())
		Expect(helpers.FileExists(command.InfiniteConfigFile())).To(BeFalse())
	})

})
