package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"

	"gopkg.in/bluesuncorp/validator.v6"

	"bitbucket.org/kiloops/api/models"
	"bitbucket.org/kiloops/api/utils"
)

var (
	client            utils.Client
	apiVersion        = "v1"
	DefaultURLProject = "https://github.com/infiniteloopsco/default-project.git"
)

//Init a http client
func Init(URL string) {
	client = utils.Client{
		&http.Client{},
		URL + "/" + apiVersion,
		"application/json",
	}
}

//PrintError e
func PrintError(e error) {
	switch t := e.(type) {
	case validator.ValidationErrors:
		fmt.Println(ErrMapString(t))
	default:
		fmt.Println("ERROR:\n")
		fmt.Println(t)
	}
}

//ErrMapString to string
func ErrMapString(errMap validator.ValidationErrors) string {
	var buffer bytes.Buffer
	buffer.WriteString("Validation errors:\n")
	for _, value := range errMap {
		buffer.WriteString(fmt.Sprintf("Field validation for '%s' failed on the field '%s'", value.Field, value.Tag))
	}
	return buffer.String()
}

//InfiniteFolder get path infinite folder
func InfiniteFolder() string {
	userPath, _ := user.Current()
	return string(userPath.HomeDir) + "/.infiniteloops"
}

//InfiniteConfigFile get path infinite folder
func InfiniteConfigFile() string {
	return InfiniteFolder() + "/config"
}

// IloopProject aasds
func IloopProject() string {
	return "iloops.project"
}

//IDLoopProjectFilePath asdas
func IDLoopProjectFileConfig(id string) string {
	return id + "/" + IloopProject()
}

//IDLoopProjectPackage asdas
func IDLoopProjectPackage(id string) string {
	return id + "/package.json"
}

//Logout the user
func LogoutFile() error {
	return os.Remove(InfiniteConfigFile())
}

func GetBodyJSON(resp *http.Response, i interface{}) {
	if jsonDataFromHTTP, err := ioutil.ReadAll(resp.Body); err == nil {
		if err := json.Unmarshal([]byte(jsonDataFromHTTP), &i); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func withUserSession(f func(*models.UserLogged) error) error {
	if user, err := loadUserSession(); err == nil {
		return f(user)
	} else {
		return err
	}
}

func loadUserSession() (*models.UserLogged, error) {
	if data, err := ioutil.ReadFile(InfiniteConfigFile()); err == nil {
		var user models.UserLogged
		json.Unmarshal(data, &user)
		return &user, nil
	} else {
		return nil, ErrNoActiveSession
	}
}

func authHeaders(user *models.UserLogged) map[string]string {
	return map[string]string{
		"AUTH_ID":    fmt.Sprintf("%d", user.ID),
		"AUTH_TOKEN": user.Token,
	}
}

func debugResponse(resp *http.Response) {
	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("*****************")
	fmt.Println(string(contents))
	fmt.Println("*****************")
}
