package downloader

import (
	"codoc/errors"
	"codoc/parser"
	"codoc/types"
	"codoc/utils"
	"io"
	"io/ioutil"
	"net/http"
)

var docLinks = map[string]string{
	"nodejs": "https://nodejs.org/api",
}

func GetDoc(docName string) (*types.Parsed, error) {
	// first check if the docs have already been downloaded
	// we need versioning for the docs
	if utils.IsAllowedDoc(docName) {
		url := docLinks[docName]
		httpResp, err := syncGet(url)

		// network error occured
		if err != nil {
			return nil, err
		}

		parsedOutput, err := parser.ParseTableOfContents(httpResp, types.Doc{DocName: docName, DocUrl: url, DocPath: ""})
		// TODO: conver to parse error
		if err != nil {
			return nil, err
		}
		return parsedOutput, err
	}
	return nil, errors.ThrowDocError("Unknown doc name", docName)
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
