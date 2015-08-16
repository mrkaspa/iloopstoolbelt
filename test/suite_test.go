package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"

	"bitbucket.org/kiloops/api/endpoint"
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ts *httptest.Server
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	fmt.Println("Suite found")
	RunSpecs(t, "Client Suite")
}

var _ = BeforeSuite(func() {
	initEnv()
	models.InitDB()
	cleanDB()
	ts = httptest.NewServer(endpoint.GetMainEngine())
	command.Init(ts.URL)
})

var _ = AfterSuite(func() {
	models.Gdb.Close()
	ts.Close()
})

func cleanDB() {
	fmt.Println("***Cleaning***")
	models.Gdb.Delete(models.UsersProjects{})
	models.Gdb.Delete(models.Execution{})
	models.Gdb.Delete(models.Project{})
	models.Gdb.Delete(models.SSH{})
	models.Gdb.Delete(models.User{})
}

func initEnv() {
	path := ".env_test"
	for i := 1; ; i++ {
		if err := godotenv.Load(path); err != nil {
			if i > 3 {
				panic("Error loading .env_test file")
			} else {
				path = "../" + path
			}
		} else {
			break
		}
	}
}
