package test

import (
	"os"

	"github.com/mrkaspa/iloopstoolbelt/command"
	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectList", func() {

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

	It("lists all the user projects", func() {
		err := command.ProjectList()
		Expect(err).To(BeNil())
	})

})
