package jenkinsapi

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AssignedLabel struct {
	Name string `json:"name" xml:"name"`
}

type BaseStruct struct {
	Class string `json:"_class" xml:"_class,attr"`
}

type IndexInfo struct {
	Class           string          `json:"_class"`
	AssignedLabels  []AssignedLabel `json:"assignedLabels" xml:"assignedLabels"`
	Mode            string          `json:"mode" xml:"mode"`
	NodeDescription string          `json:"nodeDescription" xml:"nodeDescription"`
	NodeName        string          `json:"nodeName" xml:"nodeName"`
	NumExecutors    int             `json:"numExecutors" xml:"numExecutors"`
	Description     sql.NullString  `json:"description" xml:"description"`
	Jobs            []JenkinsJob    `json:"jobs"`
	OverallLoad     interface{}     `json:"overallLoad" xml:"overallLoad"`
	PrimaryView     JenkinsView     `json:"primaryView" xml:"primaryView"`
	QuietingDown    bool            `json:"quietingDown" xml:"quietingDown"`
	SlaveAgentPort  int             `json:"slaveAgentPort" xml:"slaveAgentPort"`
	UnlabeledLoad   BaseStruct      `json:"unlabeledLoad" xml:"unlabeledLoad"`
	URL             string          `json:"url" xml:"url"`
	UseCrumbs       bool            `json:"useCrumbs" xml:"useCrumbs"`
	UseSecurity     bool            `json:"useSecurity" xml:"useSecurity"`
	Views           []JenkinsView   `json:"views"`
}

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
