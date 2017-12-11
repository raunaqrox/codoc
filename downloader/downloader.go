package downloader

import (
	"codoc/errors"
	"codoc/parser"
	"codoc/types"
	"io"
	"io/ioutil"
	"net/http"
)

// TODO: add aliases as well for each
var allowedDocs = [...]string{"nodejs", "golang"}

var docLinks = map[string]string{
	"nodejs": "https://nodejs.org/dist/latest-v9.x/docs/api/",
}

func isAllowedDoc(docName string) bool {
	for _, doc := range allowedDocs {
		if doc == docName {
			return true
		}
	}
	return false
}

func GetDoc(docName string) error {
	// first check if the docs have already been downloaded
	// we need versioning for the docs
	if isAllowedDoc(docName) {
		url := docLinks[docName]
		httpResp, err := syncGet(url)

		// network error occured
		if err != nil {
			return err
		}

		parsedOutput, err := parser.Parse(httpResp, types.Doc{DocName: docName, DocUrl: url, DocPath: ""})

		_ = parsedOutput
		_ = err

		return nil
	}
	return errors.ThrowDocError("Unknown doc name", docName)
}

func httpRespToBuffer(bodyReader io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(bodyReader)
	defer bodyReader.Close()
	if err != nil {
		// return network error with read error
		return nil, err
	}
	return body, nil
}

func syncGet(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		//TODO: return network error
		return nil, err
	}
	return resp, nil
}
