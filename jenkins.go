package jenkinsapi

import (
	"encoding/json"
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

func (j *JenkinsClient) GetJenkinsCrumb() (crumbIssue *CrumbIssue, err error) {
	req, err := http.NewRequest("GET", j.Addr+CrumbIssuerJsonURL, nil)
	if err != nil {
		return nil, err
	}
	j.SetAuth(req)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &crumbIssue)
	if err == nil {
		crumbIssue.Cookie = response.Header.Get("set-cookie")
	}
	return crumbIssue, err
}
