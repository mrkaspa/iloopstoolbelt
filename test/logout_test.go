package test

import (
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logout", func() {

	var (
		SSHPath   = "/Users/michelperez/.ssh/id_rsa.pub"
		userLogin = models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
	)

	BeforeEach(func() {
		cleanDB()
		command.CreateAccount(&userLogin, SSHPath)
		command.Login(&userLogin)
	})

	It("logouts after a login", func() {
		err := command.Logout()
		Expect(err).To(BeNil())
		Expect(helpers.FileExists(command.InfiniteConfigFile())).To(BeFalse())
	})

})
