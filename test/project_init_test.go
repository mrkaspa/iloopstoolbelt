package test

import (
	"os"

	"bitbucket.org/kiloops/api/gitadmin"
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/codeskyblue/go-sh"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectInit", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		project = defaultProject()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
	})

	It("creates a new project", func() {
		sh.Command("git", "clone", TestURLProject, project.Name).Run()
		err := command.ProjectInit(&project, true, scriptTest, cronTest)
		Expect(err).To(BeNil())
		// name := slug.Make(project.Name)
		Expect(helpers.FileExists(project.Name)).To(BeTrue())
		os.RemoveAll(project.Name)
		gitadmin.RevertAll(gitadmin.GITOLITEPATH)
	})

})
