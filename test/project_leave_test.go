package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectLeave", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		anotherUserLogin = anotherUser()
		project = defaultProject()
		command.CreateAccount(&userLogin, SSHPath)
		command.CreateAccount(&anotherUserLogin, anotherSSHPath)
		forceLogin(&userLogin)
		sh.Command("git", "clone", TestURLProject, project.Name).Run()
		command.ProjectInit(&project, true, scriptTest, cronTest)
		command.ProjectUserAdd(project.Slug, anotherUserLogin.Email)
	})

	AfterEach(func() {
		os.RemoveAll(project.Name)
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
