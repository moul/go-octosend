package octosend

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
)

var (
	APIUrl = "https://api.octosend.com/api"
)

type OctosendAPI struct {
	Token string
	Debug bool
}

func NewAPIByToken(token string) *OctosendAPI {
	return &OctosendAPI{
		Token: token,
		Debug: os.Getenv("OCTOSEND_DEBUG") == "1",
	}
}

func NewAPI(username, password string) (*OctosendAPI, error) {
	payload, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{APIUrl, "3.0/authenticate"}, "/")
	request := gorequest.New().Post(url).Send(string(payload))
	if os.Getenv("OCTOSEND_DEBUG") == "1" {
		request.SetDebug(true)
	}
	resp, body, errs := request.EndBytes()

	if len(errs) > 0 {
		return nil, printErrors(errs)
	}
	if err := httpHandleError([]int{200}, resp.StatusCode, body); err != nil {
		return nil, err
	}

	var response struct {
		Username string `json:"username"`
		ApiKey   string `json:"api-key"`
		Entity   string `json:"entity"`
	}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return NewAPIByToken(response.ApiKey), nil
}

func (a *OctosendAPI) GetRequest(path string) ([]byte, error) {
	request := gorequest.New().Get(strings.Join([]string{APIUrl, path}, "/"))
	request.Header["X-RMTA-API-Key"] = a.Token
	if a.Debug {
		request.SetDebug(true)
	}
	resp, body, errs := request.EndBytes()

	if len(errs) > 0 {
		return nil, printErrors(errs)
	}

	err := httpHandleError([]int{200}, resp.StatusCode, body)
	return body, err
}

func printErrors(errs []error) error {
	for _, err := range errs {
		logrus.Error(err)
	}
	return errors.New("Error(s) has occured")
}

func httpHandleError(goodStatusCode []int, statusCode int, body []byte) error {
	good := false
	for _, code := range goodStatusCode {
		if code == statusCode {
			good = true
		}
	}
	if !good {
		return errors.New(string(body))
	}
	return nil
}
