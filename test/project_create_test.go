package test

import (
	"os"

	"bitbucket.org/kiloops/api/gitadmin"
	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	"github.com/mrkaspa/go-helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectCreate", func() {

	BeforeEach(func() {
		userLogin = defaultUser()
		project = defaultProject()
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
	})

	It("creates a new project", func() {
		err := command.ProjectCreate(&project)
		Expect(err).To(BeNil())
		name := slug.Make(project.Name)
		Expect(helpers.FileExists(name)).To(BeTrue())
		os.RemoveAll(name)
		gitadmin.RevertAll(gitadmin.GITOLITEPATH)
	})

})
