package test

import (
	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopstoolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ForgotPassword", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		command.CreateAccount(&userLogin, SSHPath)
	})

	It("requests a new password", func() {
		err := command.ForgotPassword(&models.Email{Value: userLogin.Email})
		Expect(err).To(BeNil())
	})

})
