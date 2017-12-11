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
	ParsedDoc string `json:"doc"`
}

func init() {

}

var dat types.DocInputFormat

func Parse(toParse *http.Response, docInfo types.Doc) (*Parsed, error) {
	fileData, err := utils.ReadFile("./docsjson/" + docInfo.DocName + ".json")
	_ = err
	if err := json.Unmarshal(fileData, &dat); err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(toParse)
	_ = doc
	_ = err

	tableOfContents := doc.Find(dat.ToC)
	fmt.Println(tableOfContents.Nodes)
	_ = tableOfContents
	printNodes(tableOfContents)
	return &Parsed{
		DocInfo:   docInfo,
		ParsedDoc: "",
	}, nil
}

func printNodes(selection *goquery.Selection) {
	selection.Each(func(index int, elem *goquery.Selection) {
		name := elem.Text()
		url, _ := elem.Attr("href")
		toc := types.NewTocElem(name, url)
		fmt.Println(toc)
	})
}
