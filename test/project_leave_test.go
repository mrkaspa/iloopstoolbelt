package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectLeave", func() {

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

	It("leaves a project after assign another user as admin", func() {
		command.ProjectUserDelegate(project.Slug, anotherUserLogin.Email)
		err := command.ProjectLeave(project.Slug)
		Expect(err).To(BeNil())
	})

	It("leaves a project as an admin should fail", func() {
		err := command.ProjectLeave(project.Slug)
		Expect(err).To(BeEquivalentTo(command.ErrProjectNotLeft))
	})

})
