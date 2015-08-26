package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectUserAdd", func() {

	BeforeEach(func() {
		command.CreateAccount(&userLogin, SSHPath)
		command.CreateAccount(&anotherUserLogin, anotherSSHPath)
		forceLogin(&userLogin)
		command.ProjectCreate(&project)
	})

	AfterEach(func() {
		name := slug.Make(project.Name)
		os.RemoveAll(name)
	})

	It("adds another user to the project", func() {
		err := command.ProjectUserAdd(project.Slug, anotherUserLogin.Email)
		Expect(err).To(BeNil())
	})

	It("adds another himself to the project must fail", func() {
		err := command.ProjectUserAdd(project.Slug, userLogin.Email)
		Expect(err).NotTo(BeNil())
	})

})
