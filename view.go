package jenkinsapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

type JenkinsView struct {
	Class string `json:"_class"`
	Name  string `json:"name" xml:"name"`
	URL   string `json:"url" xml:"url"`
}

type ViewInfo struct {
	Class       string         `json:"_class"`
	Description sql.NullString `json:"description" xml:"description"`
	Jobs        []JenkinsJob   `json:"jobs"`
	Name        string         `json:"name" xml:"name"`
	URL         string         `json:"url" xml:"url"`
	Property    []UserProperty `json:"property" xml:"property"`
}

func (j *JenkinsClient) GetViewInfoJson(viewName string) (viewInfo *ViewInfo, err error) {
	req, err := http.NewRequest(
		"GET",
		j.Addr+strings.Replace(ViewInfoJsonURL, "$view", viewName, 1),
		nil,
	)
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
	err = json.Unmarshal(resp, &viewInfo)
	return viewInfo, nil
}

type ViewProperty struct {
	Class string `xml:"class,attr"`
}

type ViewConfigXML struct {
	XMLName         xml.Name
	Name            string       `xml:"name"`
	Description     string       `xml:"description"`
	FilterExecutors bool         `xml:"filterExecutors"`
	FilterQueue     bool         `xml:"filterQueue"`
	Properties      ViewProperty `xml:"properties"`
}

func (j *JenkinsClient) GetViewConfig(viewName string) (viewConfig *ViewConfigXML, err error) {
	req, err := http.NewRequest(
		"GET",
		j.Addr+strings.Replace(ViewConfigXMLURL, "$view", viewName, 1),
		nil,
	)
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
	resp2 := strings.Replace(string(resp), "1.1", "1.0", 1)
	err = xml.Unmarshal([]byte(resp2), &viewConfig)
	return viewConfig, err
}

func (j *JenkinsClient) UpdateViewConfig(viewName string, viewConfig ViewConfigXML) (err error) {
	cfgbytes, err := xml.Marshal(viewConfig)
	if err != nil {
		return err
	}
	cfg := strings.Replace(string(cfgbytes), "1.0", "1.1", 1)
	payload := bytes.NewBuffer([]byte(cfg))
	crumbIssue, err := j.GetJenkinsCrumb()
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		j.Addr+strings.Replace(ViewConfigXMLURL, "$view", viewName, 1),
		payload,
	)
	if err != nil {
		return err
	}
	j.SetAuth(req)
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	req.Header.Set(crumbIssue.CrumbRequestField, crumbIssue.Crumb)
	req.Header.Set("Cookie", crumbIssue.Cookie)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return err
}
