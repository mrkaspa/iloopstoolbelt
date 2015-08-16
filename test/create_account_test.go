package test

import (
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateAccount", func() {

	BeforeEach(func() {
	})

	It("create a new user", func() {
		userLogin := models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
		err := command.CreateAccount(&userLogin)
		Expect(err).To(BeNil())
	})

})
