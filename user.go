package jenkinsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func (j *JenkinsClient) GetUserListJson() (userList *UserList, err error) {
	req, err := http.NewRequest("GET", j.Addr+UserListJsonURL, nil)
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
	err = json.Unmarshal(resp, &userList)
	return userList, err
}

func (j *JenkinsClient) GetUserInfoJson(username string) (userList *UserInfo, err error) {
	req, err := http.NewRequest(
		"GET",
		j.Addr+strings.Replace(UserInfoJsonURL, "$username", username, 1),
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
	err = json.Unmarshal(resp, &userList)
	return userList, err
}
