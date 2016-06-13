package test

import (
	"os"

	"github.com/mrkaspa/iloopstoolbelt/command"
	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectUserRemove", func() {

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

	It("removes an user from the project", func() {
		err := command.ProjectUserRemove(project.Slug, anotherUserLogin.Email)
		Expect(err).To(BeNil())
	})

	It("removes an unkown user from the project should fail", func() {
		err := command.ProjectUserRemove(project.Slug, "demo@demo.com")
		Expect(err).To(BeEquivalentTo(command.ErrProjectOrUserNotFound))
	})

})
