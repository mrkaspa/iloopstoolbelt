package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectUserDelegate", func() {

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

	It("delegates an user as project admin", func() {
		err := command.ProjectUserDelegate(project.Slug, anotherUserLogin.Email)
		Expect(err).To(BeNil())
	})

})
