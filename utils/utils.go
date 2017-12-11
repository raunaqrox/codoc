package utils

import (
	"os/user"
)

func GetHomeFolder() string {
	usr, err := user.Current()
	if err != nil {
		panic(err) // TODO find better way to handle this
	}
	return usr.HomeDir
}
