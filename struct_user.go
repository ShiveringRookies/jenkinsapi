package jenkinsapi

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
