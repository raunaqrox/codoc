package parser

import (
	"codoc/types"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ParseDocumentationPage(toParse *http.Response) *types.DocPage {
	// topics :=
	return nil
}

// Parse the http response to html document
// apply the selections from the documentation json
// return with the documentation format
func ParseTableOfContents(toParse *http.Response, docInfo types.Doc, toc string) (*types.TableOfContents, error) {
	doc, err := goquery.NewDocumentFromResponse(toParse)

	if err != nil {
		return nil, err
	}

	tableOfContents := doc.Find(toc)

	return createDocToc(tableOfContents), nil
	// TODO use this somewhere later
	// toc.Transform()
	// return &types.Parsed{
	// 	DocInfo:   docInfo,
	// 	ParsedDoc: types.NewDocOutputFormat(toc),
	// }, nil
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
