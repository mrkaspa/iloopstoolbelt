package test

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/mrkaspa/iloopsapi/endpoint"
	"github.com/mrkaspa/iloopsapi/gitadmin"
	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopsapi/utils"
	"github.com/mrkaspa/iloopstoolbelt/command"
	gEndpoint "github.com/infiniteloopsco/guartz/endpoint"
	gModels "github.com/infiniteloopsco/guartz/models"
	gUtil "github.com/infiniteloopsco/guartz/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ts      *httptest.Server
	gServer *httptest.Server
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	fmt.Println("Suite found")
	RunSpecs(t, "Client Suite")
}

var _ = BeforeSuite(func() {
	initEnv()
	utils.InitLogTest()
	utils.InitEmail()
	gUtil.InitLogTest()
	models.InitDB()
	gModels.InitDB()
	gitadmin.InitVars()
	gitadmin.InitGitAdmin()
	gModels.InitCron()
	cleanDB()
	ts = httptest.NewServer(endpoint.GetMainEngine())
	gServer = httptest.NewServer(gEndpoint.GetMainEngine())
	gURL, _ := url.Parse(gServer.URL)
	os.Setenv("GUARTZ_HOST", gURL.Host)
	models.InitGuartzClient()
	command.Init(ts.URL)
	cleaner()
})

var _ = AfterSuite(func() {
	models.Gdb.Close()
	gModels.Gdb.Close()
	ts.Close()
	gServer.Close()
	gitadmin.FinishGitAdmin()
})

var _ = BeforeEach(func() {
	cleaner()
})

func cleaner() {
	cleanDB()
	gitadmin.RevertAll(gitadmin.GITOLITEPATH)
	os.Remove(command.InfiniteConfigFile())
	os.RemoveAll(command.InfiniteFolder())
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
	gModels.Gdb.Delete(gModels.Execution{})
	gModels.Gdb.Delete(gModels.Task{})
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
