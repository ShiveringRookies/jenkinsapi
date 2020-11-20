package jenkinsapi

type AssignedLabel struct {
	Name string `json:"name" xml:"name"`
}

type BaseStruct struct {
	Class string `json:"_class"`
}

type IndexInfo struct {
	Class           string          `json:"_class"`
	AssignedLabels  []AssignedLabel `json:"assignedLabels" xml:"assignedLabels"`
	Mode            string          `json:"mode" xml:"mode"`
	NodeDescription string          `json:"nodeDescription" xml:"nodeDescription"`
	NodeName        string          `json:"nodeName" xml:"nodeName"`
	NumExecutors    int             `json:"numExecutors" xml:"numExecutors"`
	Description     interface{}     `json:"description" xml:"description"`
	Jobs            []JenkinsJob    `json:"jobs"`
	OverallLoad     interface{}     `json:"overallLoad" xml:"overallLoad"`
	PrimaryView     JenkinsView     `json:"primaryView" xml:"primaryView"`
	QuietingDown    bool            `json:"quietingDown" xml:"quietingDown"`
	SlaveAgentPort  int             `json:"slaveAgentPort" xml:"slaveAgentPort"`
	UnlabeledLoad   BaseStruct      `json:"unlabeledLoad" xml:"unlabeledLoad"`
	URL             string          `json:"url" xml:"url"`
	UseCrumbs       bool            `json:"useCrumbs" xml:"useCrumbs"`
	UseSecurity     bool            `json:"useSecurity" xml:"useSecurity"`
	Views           []JenkinsView   `json:"views"`
}
