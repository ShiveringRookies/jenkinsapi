package jenkinsapi

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
