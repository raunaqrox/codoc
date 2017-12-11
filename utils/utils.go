package utils

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// TODO make it platform agnostic
// get home folder of user
func GetHomeFolder() string {
	usr, err := user.Current()
	if err != nil {
		panic(err) // TODO find better way to handle this
	}
	return usr.HomeDir
}

// get codoc folder path
func GetCodocFolder() string {
	homeFolder := GetHomeFolder()
	// TODO Solve the cyclic issue due to importing config here
	return filepath.Join(homeFolder, ".codoc") // cyclic dependency if called from "codoc/config" package
}

// find if codoc folder is created
func FolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	// no error hence it exists
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err

}

func ListFilesInFolder(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}
	return files
}

func GetDocPath(docName string) string {
	return filepath.Join(GetCodocFolder(), docName+".json")
}

func IsDocInstalled(docName string) bool {
	exists, err := FolderExists(GetDocPath(docName))
	if err != nil {
		panic(err)
	}
	return exists
}
