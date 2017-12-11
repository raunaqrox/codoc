package config

import (
	"os/user"
)

var Config = map[string]string{
	"homeFolder": getHomeFolder(),
}

func getHomeFolder() string {
	usr, err := user.Current()
	if err != nil {
		panic(err) // TODO find better way to handle this
	}
	return usr.HomeDir
}
