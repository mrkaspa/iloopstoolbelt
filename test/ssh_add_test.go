package test

import (
	"bitbucket.org/kiloops/toolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSHAdd", func() {

	BeforeEach(func() {
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
	})

	It("uploads a new SSH", func() {
		err := command.SSHAdd("Another ssh", anotherSSHPath)
		Expect(err).To(BeNil())
	})

})
