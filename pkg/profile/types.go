package profile

type Profile struct {
	Servers        []Server  `json:"servers"`
	Users          []User    `json:"users"`
	Contexts       []Context `json:"contexts"`
	CurrentContext string    `json:"current-context,omitempty"`
}

type Context struct {
	Name    string `json:"name,omitempty"`
	User    string `json:"user,omitempty"`
	Server  string `json:"server,omitempty"`
	Channel string `json:"channel,omitempty"`
}

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Server struct {
	Url  string `json:"url,omitempty"`
}
