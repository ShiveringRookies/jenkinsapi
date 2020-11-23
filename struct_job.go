package jenkinsapi

type JenkinsJob struct {
	Class string `json:"class"`
	Name  string `json:"name" xml:"name"`
	URL   string `json:"url" xml:"url"`
	Color string `json:"color" xml:"color"`
}
