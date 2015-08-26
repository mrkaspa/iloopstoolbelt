package test

import (
	"os"

	"bitbucket.org/kiloops/toolbelt/command"
	"github.com/gosimple/slug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectDelete", func() {

	BeforeEach(func() {
		command.CreateAccount(&userLogin, SSHPath)
		forceLogin(&userLogin)
		command.ProjectCreate(&project)
	})

	AfterEach(func() {
		name := slug.Make(project.Name)
		os.RemoveAll(name)
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
