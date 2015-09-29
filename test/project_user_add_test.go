package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectUserAdd", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		anotherUserLogin = anotherUser()
		project = defaultProject()
		command.CreateAccount(&userLogin, SSHPath)
		command.CreateAccount(&anotherUserLogin, anotherSSHPath)
		forceLogin(&userLogin)
		sh.Command("git", "clone", TestURLProject, project.Name).Run()
		command.ProjectInit(&project, true, scriptTest, cronTest)
	})

	AfterEach(func() {
		os.RemoveAll(project.Name)
	})

	It("adds another user to the project", func() {
		err := command.ProjectUserAdd(project.Slug, anotherUserLogin.Email)
		Expect(err).To(BeNil())
	})

	It("adds himself again to the project must fail", func() {
		err := command.ProjectUserAdd(project.Slug, userLogin.Email)
		Expect(err).NotTo(BeNil())
	})

})
