package test

import (
	"github.com/mrkaspa/iloopstoolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSHAdd", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
	})

	It("uploads a new SSH", func() {
		err := command.SSHAdd("Another ssh", anotherSSHPath)
		Expect(err).To(BeNil())
	})

})
