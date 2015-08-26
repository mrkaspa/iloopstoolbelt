package test

import (
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/bluesuncorp/validator.v6"
)

var _ = Describe("CreateAccount", func() {

	It("create a new user", func() {
		userLogin = models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
		err := command.CreateAccount(&userLogin, SSHPath)
		Expect(err).To(BeNil())
	})

	It("create a user with a bad email", func() {
		userBadLogin := models.UserLogin{Email: "michel.ingesoft", Password: "h1h1h1h1h1h1"}
		err := command.CreateAccount(&userBadLogin, SSHPath)
		Expect(err).NotTo(BeNil())
		errMap := err.(validator.ValidationErrors)
		Expect(errMap["UserLogin.Email"]).NotTo(BeNil())
	})

	Context("after creating an user", func() {

		BeforeEach(func() {
			command.CreateAccount(&userLogin, SSHPath)
		})

		It("create a new user with the same email", func() {
			err := command.CreateAccount(&userLogin, SSHPath)
			Expect(err).NotTo(BeNil())
		})

	})

})
