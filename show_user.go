package jenkinsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type User struct {
	AbsoluteURL string `json:"absoluteUrl" xml:"absoluteUrl"`
	FullName    string `json:"fullName" xml:"fullName"`
}

type UserSimpleInfo struct {
	LastChange interface{} `json:"lastChange" xml:"-"`
	Project    interface{} `json:"project" xml:"-"`
	User       User        `json:"user" xml:"user"`
}

type UserList struct {
	Class string           `json:"_class"`
	Users []UserSimpleInfo `json:"users"`
}

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

type UserProperty struct {
	Class             string `json:"_class"`
	Address           string `json:"address,omitempty"`
	InsensitiveSearch bool   `json:"insensitiveSearch,omitempty"`
}

type UserInfo struct {
	Class       string         `json:"_class"`
	AbsoluteURL string         `json:"absoluteUrl" xml:"absoluteUrl"`
	ID          string         `json:"id" xml:"id"`
	Description string         `json:"description" xml:"description"`
	FullName    string         `json:"fullName" xml:"fullName"`
	Property    []UserProperty `json:"property"`
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
