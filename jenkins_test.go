package jenkinsapi

import "testing"

func Test_InitJenkinsClient(t *testing.T) {
	j := JenkinsClient{
		Addr: "http://127.0.0.1:8080",
		BasicAuth: BasicAuth{
			UserName: "xiaoboya",
			Password: "xby951111",
		},
		BearerAuth: "",
	}
	err := j.Connect()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
