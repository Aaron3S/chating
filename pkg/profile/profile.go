package profile

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const profilePath = "~/.chat/profile"
const fileName = "config"

var errProfileNotFound = fmt.Errorf("can not find profile file in %s", profilePath)

func View() (Profile, error) {
	filepath := path.Join(profileRealPath(), fileName)

	if !fileutil.Exist(filepath) {
		return Profile{}, errProfileNotFound
	}
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Profile{}, err
	}
	var p Profile
	if err := yaml.Unmarshal(bs, &p); err != nil {
		return Profile{}, err
	}
	return p, nil
}

func GetContext(contextName string) (Context, error) {
	p, err := prepareProfile()
	if err != nil {
		return Context{}, err
	}
	for i := range p.Contexts {
		if p.Contexts[i].Name == contextName {
			return p.Contexts[i], nil
		}
	}
	return Context{}, nil
}

func AddUser(user *User) error {
	p, err := prepareProfile()
	if err != nil {
		return err
	}
	for i := range p.Users {
		if p.Users[i].Name == user.Name {
			return fmt.Errorf("user %s already exists", user.Name)
		}
	}
	p.Users = append(p.Users, *user)
	return saveProfile(p)
}

func AddServer(server *Server) error {
	p, err := prepareProfile()
	if err != nil {
		return err
	}
	for i := range p.Servers {
		if p.Servers[i].Url == server.Url {
			return fmt.Errorf("server %s already exists", server.Url)
		}
	}
	p.Servers = append(p.Servers, *server)
	fmt.Println(1)
	return saveProfile(p)
}

func AddContext(ctx *Context) error {
	p, err := prepareProfile()
	if err != nil {
		return err
	}
	var userExists, serverExists bool
	for i := range p.Users {
		if ctx.User == p.Users[i].Name {
			userExists = true
		}
	}
	for i := range p.Servers {
		if ctx.Server == p.Servers[i].Url {
			serverExists = true
		}
	}
	if !userExists {
		return fmt.Errorf("can not find user %s in profile", ctx.User)
	}
	if !serverExists {
		return fmt.Errorf("can not find server %s in profile", ctx.Server)
	}
	for i := range p.Contexts {
		if p.Contexts[i].Name == ctx.Name {
			return fmt.Errorf("context %s already exists", ctx.Name)
		}
	}
	p.Contexts = append(p.Contexts, *ctx)
	if len(p.Contexts) == 1 {
		p.CurrentContext = p.Contexts[0].Name
	}
	return saveProfile(p)
}

func prepareProfile() (Profile, error) {
	//ensure path exists
	if !fileutil.Exist(profileRealPath()) {
		if err := fileutil.CreateDirAll(profileRealPath()); err != nil {
			return Profile{}, err
		}
	}
	filepath := path.Join(profileRealPath(), fileName)
	var p Profile
	var err error
	if fileutil.Exist(filepath) {
		p, err = View()
		if err != nil {
			return Profile{}, err
		}
	} else {
		f, err := os.Create(filepath)
		if err != nil {
			return Profile{}, err
		}
		defer f.Close()
	}
	return p, nil
}
func saveProfile(p Profile) error {
	bs, err := yaml.Marshal(&p)
	if err != nil {
		return err
	}
	filepath := path.Join(profileRealPath(), fileName)
	return ioutil.WriteFile(filepath, bs, 0755)
}

func profileRealPath() string {
	hd, _ := homedir.Dir()
	s := strings.Replace(profilePath, "~", hd, -1)
	return s
}
