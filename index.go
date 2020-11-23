package jenkinsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (j *JenkinsClient) Index() (indexInfo *IndexInfo, err error) {
	req, err := http.NewRequest("GET", j.Addr+IndexJsonURL, nil)
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
	err = json.Unmarshal(resp, &indexInfo)
	return indexInfo, err
}

type CrumbIssue struct {
	Class             string `json:"_class"`
	Crumb             string `json:"crumb"`
	CrumbRequestField string `json:"crumbRequestField"`
	Cookie            string `json:"cookie,omitempty"`
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
