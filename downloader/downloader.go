package downloader

import (
	"codoc/errors"
	"codoc/parser"
	"codoc/types"
	"codoc/utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var docLinks = map[string]string{
	"nodejs": "https://nodejs.org/api/",
}

// fetch all the http responses for each page
func getDocPages(toc *types.TableOfContents, baseUrl string, topicSelector string) ([]*http.Response, error) {
	result := make([]*http.Response, len(toc.Toc))
	for i, topic := range toc.Toc {
		resolvedUrl, err := utils.ResolveUrl(baseUrl, topic.Link)
		if err != nil {
			return nil, err
		}
		resp, err := syncGet(resolvedUrl.String())
		if err != nil {
			return nil, err
		}
		parser.ParseDocumentationPage(resp, topicSelector)
		result[i] = resp
	}
	return result, nil
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

		jsonStruct, err := utils.ReadDocJson(filepath.Join("./docsjson/", docName+".json"))
		_ = err

		tableOfContents, err := parser.ParseTableOfContents(httpResp, types.Doc{DocName: docName, DocUrl: url, DocPath: ""}, jsonStruct.Toc)
		docPages, err := getDocPages(tableOfContents, url, jsonStruct.Topic)
		_ = err
		fmt.Println(docPages)
		s, err := httpRespToBuffer(docPages[0].Body)
		_ = err
		fmt.Println(string(s))
		// TODO: convert to parse error
		if err != nil {
			return nil, err
		}

		return nil, err
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
