package downloader

import (
	"codoc/errors"
	"codoc/parser"
	"codoc/types"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// TODO: add aliases as well for each
var allowedDocs = [...]string{"nodejs", "golang"}

var docLinks = map[string]string{
	"nodejs": "https://nodejs.org/en/docs/",
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

		bodyBuffer, err := httpRespToBuffer(httpResp.Body)

		if err != nil {
			return err
		}

		parsedOutput := parser.Parse(string(bodyBuffer), types.Doc{DocName: docName, DocUrl: url, DocPath: ""})

		fmt.Printf("parsed output %s", parsedOutput)

		fmt.Printf("%v", string(bodyBuffer))
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

// func asyncHttpGet(url string) *HttpResponse {
// 	ch := make(chan *HttpResponse)
// 	go func(url string) {
// 		fmt.Printf("fetching %s \n", url)
// 		resp, err := http.Get(url)
// 		defer resp.Body.Close()
// 		ch <- &HttpResponse{url, resp, err}
// 	}(url)

// 	for {
// 		select {
// 		case r := <-ch:
// 			fmt.Printf("\n%s was fetched \n", r.url)
// 			return r
// 		case <-time.After(50 * time.Millisecond):
// 			fmt.Printf(".")
// 		}
// 	}
// }
