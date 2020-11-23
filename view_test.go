package jenkinsapi

import (
	"encoding/xml"
	"testing"
)

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

func TestJenkinsClient_GetViewConfig(t *testing.T) {
	j := JenkinsClient{
		Addr: "http://127.0.0.1:8080",
		BasicAuth: BasicAuth{
			UserName: "xiaoboya",
			Password: "xby951111",
		},
		BearerAuth: "",
	}
	viewInfo, err := j.GetViewConfig("TestMyView")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(viewInfo)
	}
}

func TestJenkinsClient_UpdateViewConfig(t *testing.T) {
	j := JenkinsClient{
		Addr: "http://127.0.0.1:8080",
		BasicAuth: BasicAuth{
			UserName: "xiaoboya",
			Password: "xby951111",
		},
		BearerAuth: "",
	}
	err := j.UpdateViewConfig(
		"TestMyView",
		ViewConfigXML{
			xml.Name{Local: "hudson.model.MyView"},
			"TestMyView",
			"测试用的视图",
			true, false,
			ViewProperty{Class: "hudson.model.View$PropertyList"},
		},
	)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
