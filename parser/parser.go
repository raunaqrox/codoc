package parser

import (
	"codoc/types"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Parse http response to DocPages
func ParseDocumentationPage(toParse *http.Response, topicSelector string) (*types.DocPage, error) {
	doc, err := goquery.NewDocumentFromResponse(toParse)

	if err != nil {
		return nil, err
	}

	topics := doc.Find(topicSelector)
	nodeDoc := &types.Nodejs{topics}
	_ = nodeDoc

	section := nodeDoc.Transform()
	_ = section
	// fmt.Printf("section %v", section)

	return nil, nil
}

// Parse the http response to Table of contents of documentation
// Apply the selections from the documentation json
func ParseTableOfContents(toParse *http.Response, docInfo types.Doc, tocSelector string) (*types.TableOfContents, error) {
	doc, err := goquery.NewDocumentFromResponse(toParse)

	if err != nil {
		return nil, err
	}

	fmt.Println(tocSelector)
	tableOfContents := doc.Find(tocSelector)

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
