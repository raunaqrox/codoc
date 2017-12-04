package downloader

import "codoc/errors"

// TODO: add aliases as well for each
var allowedDocs = [...]string{"nodejs", "golang"}

func isAllowedDoc(docName string) bool {
	for _, doc := range allowedDocs {
		if doc == docName {
			return true
		}
	}
	return false
}

type Doc struct {
	docName string
	docUrl  string
	docPath string
}

func GetDoc(docName string) error {
	// first check if the docs have already been downloaded
	// we need versioning for the docs
	if isAllowedDoc(docName) {

		return nil
	}
	return errors.ThrowDocError("Unknown doc name", docName)
}
