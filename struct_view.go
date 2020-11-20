package jenkinsapi

import "database/sql"

type JenkinsView struct {
	Class string `json:"_class"`
	Name  string `json:"name" xml:"name"`
	URL   string `json:"url" xml:"url"`
}

type ViewInfo struct {
	Class       string         `json:"_class"`
	Description sql.NullString `json:"description" xml:"description"`
	Jobs        []JenkinsJob   `json:"jobs"`
	Name        string         `json:"name" xml:"name"`
	URL         string         `json:"url" xml:"url"`
	Property    []UserProperty `json:"property" xml:"property"`
}
