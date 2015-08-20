package test

import (
	"os"

	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectCreate", func() {

	var (
		SSHPath   = "/Users/michelperez/.ssh/id_rsa.pub"
		userLogin = models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
		project   = models.Project{Name: "demo"}
	)

	BeforeEach(func() {
		cleanDB()
		command.CreateAccount(&userLogin, SSHPath)
		command.Login(&userLogin)
	})

	It("creates a new project", func() {
		err := command.ProjectCreate(&project)
		Expect(err).To(BeNil())
		name := slug.Make(project.Name)
		Expect(helpers.FileExists(name)).To(BeTrue())
		os.RemoveAll(name)
	})

})
