package test

import (
	"github.com/mrkaspa/iloopstoolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSHDelete", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
		command.SSHAdd("another-ssh", anotherSSHPath)
	})

	It("delete the SSH", func() {
		err := command.SSHDelete("another-ssh")
		Expect(err).To(BeNil())
	})

})
