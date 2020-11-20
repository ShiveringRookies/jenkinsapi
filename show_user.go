package jenkinsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	AbsoluteURL string `json:"absoluteUrl" xml:"absoluteUrl"`
	FullName    string `json:"fullName" xml:"fullName"`
}

type UserInfo struct {
	LastChange interface{} `json:"lastChange" xml:"-"`
	Project    interface{} `json:"project" xml:"-"`
	User       User        `json:"user" xml:"user"`
}

type UserList struct {
	Class string     `json:"_class"`
	Users []UserInfo `json:"users"`
}

func (j *JenkinsClient) GetUserListJson() (userList *UserList, err error) {
	req, err := http.NewRequest("GET", j.Addr+"/asynchPeople/api/json", nil)
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
