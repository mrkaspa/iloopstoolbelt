package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectList", func() {

	BeforeEach(func() {
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
		command.ProjectCreate(&project)
	})

	AfterEach(func() {
		name := slug.Make(project.Name)
		os.RemoveAll(name)
	})

	It("lists all the user projects", func() {
		err := command.ProjectList()
		Expect(err).To(BeNil())
	})

})
