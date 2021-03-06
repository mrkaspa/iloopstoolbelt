package command

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"

	"github.com/codegangsta/cli"
	"github.com/howeyc/gopass"

	"gopkg.in/bluesuncorp/validator.v6"

	"github.com/mrkaspa/iloopsapi/models"
	"github.com/mrkaspa/iloopsapi/utils"
)

var (
	client             utils.Client
	apiVersion         = "v1"
	ValidationMessages = map[string]string{
		"email":    "the %s has an invalid format",
		"required": "the %s is required",
	}
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
		fmt.Println("ERROR:")
		fmt.Println(t)
	}
}

//ErrMapString to string
func ErrMapString(errMap validator.ValidationErrors) string {
	var buffer bytes.Buffer
	buffer.WriteString("Validation errors:\n")
	for _, value := range errMap {
		var msg string
		if template, ok := ValidationMessages[value.Tag]; ok {
			msg = fmt.Sprintf(template, value.Field)
		} else {
			msg = fmt.Sprintf("Field validation for '%s' failed on the field '%s'", value.Tag, value.Field)
		}
		buffer.WriteString(msg)
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

func withUserSession(f func(*models.UserLogged) error) error {
	if user, err := loadUserSession(); err == nil {
		return f(user)
	} else {
		return err
	}
}

func validateArgAt(args cli.Args, pos int) error {
	if len(args) < (pos+1) && args[pos] != "" {
		return ErrArgNotFound
	}
	return nil
}

//LoginFile configuration file
func loginFile(user *models.UserLogged) error {
	if err := os.Mkdir(InfiniteFolder(), os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	if _, err := os.Create(InfiniteConfigFile()); err != nil {
		return err
	}
	authJSON, _ := json.Marshal(user)
	secureContent, err := cypher(authJSON)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(InfiniteConfigFile(), secureContent, os.ModePerm)
}

func loadUserSession() (*models.UserLogged, error) {
	if secureContent, err := ioutil.ReadFile(InfiniteConfigFile()); err == nil {
		authJSON, err := decypher(secureContent)
		if err != nil {
			return nil, err
		}
		var user models.UserLogged
		json.Unmarshal(authJSON, &user)
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

func readLine(prompt string) string {
	var in string
	for in == "" {
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		in, _ = reader.ReadString('\n')
	}
	sz := len(in)
	if sz > 0 && in[sz-1] == '\n' {
		in = in[:sz-1]
	}
	return in
}

func readPassword(prompt string) string {
	var password string
	for password == "" {
		fmt.Println(prompt)
		data := gopass.GetPasswd()
		password = string(data)
	}
	return password
}

func debugResponse(resp *http.Response) {
	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("*****************")
	fmt.Println(string(contents))
	fmt.Println("*****************")
}

func shortString(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
