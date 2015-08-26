package test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"bitbucket.org/kiloops/api/endpoint"
	"bitbucket.org/kiloops/api/gitadmin"
	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/toolbelt/command"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	gitadmin.InitVars()
	gitadmin.InitGitAdmin()
	ts = httptest.NewServer(endpoint.GetMainEngine())
	command.Init(ts.URL)
	cleaner()
})

var _ = AfterSuite(func() {
	models.Gdb.Close()
	ts.Close()
	gitadmin.FinishGitAdmin()
})

var _ = BeforeEach(func() {
	time.Sleep(2 * time.Second)
	cleaner()
})

func cleaner() {
	cleanDB()
	gitadmin.RevertAll(gitadmin.GITOLITEPATH)
	os.RemoveAll(command.InfiniteConfigFile())
}

func forceLogin(userLogin *models.UserLogin) {
	if err := command.Login(userLogin); err != nil {
		fmt.Println(err)
		panic("Login Error")
	}
}

func cleanDB() {
	fmt.Println("***Cleaning***")
	models.Gdb.Delete(models.UsersProjects{})
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
