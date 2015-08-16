package command

import (
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/kiloops/api/models"
)

var (
	client     Client
	apiVersion = "v1"
)

func Init(URL string) {
	client = Client{
		&http.Client{},
		URL + "/" + apiVersion,
		"application/json",
	}
}

//Client for http requests
type Client struct {
	*http.Client
	baseURL     string
	contentType string
}

func (c Client) CallRequest(method string, path string, reader io.Reader) (*http.Response, error) {
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

func (c Client) CallRequestWithHeaders(method string, path string, reader io.Reader, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest(method, c.baseURL+path, reader)
	req.Header.Set("Content-Type", c.contentType)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	return c.Do(req)
}

func authHeaders(user models.User) map[string]string {
	return map[string]string{
		"AUTH_ID":    fmt.Sprintf("%d", user.ID),
		"AUTH_TOKEN": user.Token,
	}
}
