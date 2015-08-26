package test

import (
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logout", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		command.CreateAccount(&userLogin, SSHPath)
		command.Login(&userLogin)
	})

	It("logouts after a login", func() {
		err := command.Logout()
		Expect(err).To(BeNil())
		Expect(helpers.FileExists(command.InfiniteConfigFile())).To(BeFalse())
	})

})
