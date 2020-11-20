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
