package jenkinsapi

type JenkinsView struct {
	Class string `json:"class"`
	Name  string `json:"name" xml:"name"`
	URL   string `json:"url" xml:"url"`
}
