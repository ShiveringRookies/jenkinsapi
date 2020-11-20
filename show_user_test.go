package jenkinsapi

import "testing"

func TestJenkinsClient_GetUserListJson(t *testing.T) {
	j := JenkinsClient{
		Addr: "http://127.0.0.1:8080",
		BasicAuth: BasicAuth{
			UserName: "xiaoboya",
			Password: "xby951111",
		},
		BearerAuth: "",
	}
	userList, err := j.GetUserListJson()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(userList)
	}
}
