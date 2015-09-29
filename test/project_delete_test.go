package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectDelete", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		project = defaultProject()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
		sh.Command("git", "clone", TestURLProject, project.Name).Run()
		command.ProjectInit(&project, true, scriptTest, cronTest)
	})

	AfterEach(func() {
		os.RemoveAll(project.Name)
	})

	It("deletes an existing project", func() {
		err := command.ProjectDelete(project.Slug)
		Expect(err).To(BeNil())
	})

	It("deletes an unexisting project so should fail", func() {
		err := command.ProjectDelete("demo-un")
		Expect(err).To(BeEquivalentTo(command.ErrProjectNotFound))
	})

})
