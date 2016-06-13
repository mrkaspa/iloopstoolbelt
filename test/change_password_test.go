package test

import (
	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopstoolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChangePassword", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		command.CreateAccount(&userLogin, SSHPath)
	})

	It("changes the password with a wrong token", func() {
		err := command.ChangePassword(&models.ChangePassword{Token: "wrong-token", Password: "joka123"})
		Expect(err).NotTo(BeNil())
	})

})
