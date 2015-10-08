package test

import (
	"bitbucket.org/kiloops/toolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSHList", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
	})

	It("lists all the SSH", func() {
		err := command.SSHList()
		Expect(err).To(BeNil())
	})

})
