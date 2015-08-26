package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectDeploy", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		project = defaultProject()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
		command.ProjectCreate(&project)
	})

	AfterEach(func() {
		name := slug.Make(project.Name)
		os.RemoveAll(name)
	})

	It("deploys the current project", func() {
		err := command.ProjectDeploy()
		Expect(err).To(BeNil())
	})

})
