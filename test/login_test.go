package test

import (
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
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
