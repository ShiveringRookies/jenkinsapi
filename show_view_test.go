package jenkinsapi

import "testing"

func TestJenkinsClient_GetViewInfo(t *testing.T) {
	j := JenkinsClient{
		Addr: "http://127.0.0.1:8080",
		BasicAuth: BasicAuth{
			UserName: "xiaoboya",
			Password: "xby951111",
		},
		BearerAuth: "",
	}
	viewInfo, err := j.GetViewInfoJson("TestMyView")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(viewInfo)
	}
}
