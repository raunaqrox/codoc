package parser

import (
	"codoc/types"
	"codoc/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Parsed struct {
	DocInfo   types.Doc
	ParsedDoc *types.DocOutputFormat
}

func init() {

}

var dat types.DocInputFormat

// Parse the http response to html document
// apply the selections from the documentation json
// return with the documentation format
func Parse(toParse *http.Response, docInfo types.Doc) (*Parsed, error) {
	fileData, err := utils.ReadFile("./docsjson/" + docInfo.DocName + ".json")
	_ = err
	if err := json.Unmarshal(fileData, &dat); err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(toParse)
	_ = doc
	_ = err

	tableOfContents := doc.Find(dat.Toc)
	toc := createDocToc(tableOfContents)
	fmt.Println(toc)
	return &Parsed{
		DocInfo:   docInfo,
		ParsedDoc: types.NewDocOutputFormat(toc),
	}, nil
}

// create the table of contents of the documentation
func createDocToc(selection *goquery.Selection) []*types.TocElem {
	tableOfContents := make([]*types.TocElem, len(selection.Nodes))
	selection.Each(func(index int, elem *goquery.Selection) {
		name := elem.Text()
		// resolve the href with the correct url
		// fetch the current url and add it to it
		url, _ := elem.Attr("href")
		toc := types.NewTocElem(name, url)
		tableOfContents[index] = toc
	})
	return tableOfContents
}
