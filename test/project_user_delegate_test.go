package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectUserDelegate", func() {

	BeforeEach(func() {
		cleanDB()
		command.CreateAccount(&userLogin, SSHPath)
		command.CreateAccount(&anotherUserLogin, anotherSSHPath)
		command.Login(&userLogin)
		command.ProjectCreate(&project)
		command.ProjectUserAdd(project.Slug, anotherUserLogin.Email)
	})

	AfterEach(func() {
		name := slug.Make(project.Name)
		os.RemoveAll(name)
	})

	It("delegates an user as project admin", func() {
		err := command.ProjectUserDelegate(project.Slug, anotherUserLogin.Email)
		Expect(err).To(BeNil())
	})

})
