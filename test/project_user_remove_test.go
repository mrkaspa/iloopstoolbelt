package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectUserRemove", func() {

	BeforeEach(func() {
		cleanDB()
		command.CreateAccount(&userLogin, SSHPath)
		command.CreateAccount(&anotherUserLogin, anotherSSHPath)
		command.Login(&userLogin)
		command.ProjectCreate(&project)
		command.ProjectUserAdd(project.Slug, anotherUserLogin.Email)
	})

	AfterEach(func() {
		name := slug.Make(project.Name)
		os.RemoveAll(name)
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
