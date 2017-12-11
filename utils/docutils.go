package utils

import (
	"codoc/fs"
	"codoc/types"
	"encoding/json"
	"path/filepath"
	"strings"
)

// read this too from the config maybe
var allowedDocs = [...]string{"nodejs", "golang"}

func IsAllowedDoc(docName string) bool {
	for _, doc := range allowedDocs {
		if doc == docName {
			return true
		}
	}
	return false
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

func GetDocList(path string) []string {
	files := ListFilesInFolder(path)
	docs := make([]string, len(files))
	for i, f := range files {
		docs[i] = strings.Split(f.Name(), ".")[0]
	}
	return docs
}

// instead of getting the table of contents first get more fields in the json
// extract the json in a top level doc type
func GetDoc(docName string) (*types.TableOfContents, error) {
	var docOutput types.TableOfContents
	fileData, err := fs.ReadFile(filepath.Join(GetCodocFolder(), docName+".json"))
	if err = json.Unmarshal(fileData, &docOutput); err != nil {
		return nil, err
	}

	return &docOutput, nil
}
