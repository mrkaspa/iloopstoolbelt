package command

import (
	"bytes"
	"fmt"
	"net/http"

	"gopkg.in/bluesuncorp/validator.v6"

	"bitbucket.org/kiloops/api/utils"
)

var (
	client     utils.Client
	apiVersion = "v1"
)

func Init(URL string) {
	client = utils.Client{
		&http.Client{},
		URL + "/" + apiVersion,
		"application/json",
	}
}

func PrintError(e error) {
	switch t := e.(type) {
	case validator.ValidationErrors:
		fmt.Println(ErrMapString(t))
	default:
		fmt.Println("ERROR:\n")
		fmt.Println(t)
	}
}

func ErrMapString(errMap validator.ValidationErrors) string {
	var buffer bytes.Buffer
	buffer.WriteString("Validation errors:\n")
	for _, value := range errMap {
		buffer.WriteString(fmt.Sprintf("Field validation for '%s' failed on the field '%s'", value.Field, value.Tag))
	}
	return buffer.String()
}

// func authHeaders(user models.User) map[string]string {
// 	return map[string]string{
// 		"AUTH_ID":    fmt.Sprintf("%d", user.ID),
// 		"AUTH_TOKEN": user.Token,
// 	}
// }
