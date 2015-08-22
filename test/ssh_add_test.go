package test

import (
	"bitbucket.org/kiloops/toolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSHAdd", func() {

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
