package parser

import (
	"codoc/types"
	"codoc/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Parse the http response to html document
// apply the selections from the documentation json
// return with the documentation format
func Parse(toParse *http.Response, docInfo types.Doc) (*types.Parsed, error) {
	var dat types.DocInputFormat
	fileData, err := utils.ReadFile("./docsjson/" + docInfo.DocName + ".json")
	_ = err
	if err := json.Unmarshal(fileData, &dat); err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(toParse)

	if err != nil {
		return nil, err
	}

	tableOfContents := doc.Find(dat.Toc)
	toc := createDocToc(tableOfContents)
	fmt.Println(toc.Transform())
	return &types.Parsed{
		DocInfo:   docInfo,
		ParsedDoc: types.NewDocOutputFormat(toc),
	}, nil
}

// create the table of contents of the documentation
func createDocToc(selection *goquery.Selection) *types.TableOfContents {
	tableOfContents := make([]*types.TocElem, len(selection.Nodes))
	selection.Each(func(index int, elem *goquery.Selection) {
		name := elem.Text()
		// resolve the href with the correct url
		// fetch the current url and add it to it
		// only if the attr is found else ignore
		if url, ok := elem.Attr("href"); ok {
			toc := types.NewTocElem(name, url)
			tableOfContents[index] = toc
		}
	})
	return &types.TableOfContents{tableOfContents}
}
