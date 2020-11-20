package jenkinsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

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
