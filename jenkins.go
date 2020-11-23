package jenkinsapi

import (
	"io/ioutil"
	"net/http"
)

type BasicAuth struct {
	UserName string
	Password string
}

type JenkinsClient struct {
	Addr       string
	BasicAuth  BasicAuth
	BearerAuth string
}

func (j *JenkinsClient) Connect() (err error) {
	req, err := http.NewRequest("GET", j.Addr, nil)
	if err != nil {
		return err
	}
	j.SetAuth(req)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	return nil
}

func (j *JenkinsClient) SetAuth(req *http.Request) {
	switch j.BearerAuth {
	case "":
		req.SetBasicAuth(j.BasicAuth.UserName, j.BasicAuth.Password)
	default:
	}
}
