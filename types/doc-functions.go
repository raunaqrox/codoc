package types

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// temporary way for now to handle transformations
// for each key with selector to help parse the page
// we get the element and call the corresponding function written here
/*
var docKeyHandlers = map[string]interface{}{
	"table_of_contents": func(tableOfContents []*TocElem) []*TocElem {
		return tableOfContents
	},
}
*/

func NewDocPage(toc *TableOfContents, meta interface{}) *DocPage {
	return &DocPage{
		LocalToc: toc,
		MetaData: meta,
	}
}

func NewExplanation(explanation string, example string) *Explanation {
	return &Explanation{
		explanation,
		example,
	}
}

func NewSection(topic string, explanation []Explanation) Section {
	return Section{
		topic,
		explanation,
	}
}

func NewTocElem(name, link string) *TocElem {
	return &TocElem{
		name,
		link,
	}
}

func NewDocOutputFormat(todoElems *TableOfContents) *DocOutputFormat {
	return &DocOutputFormat{
		todoElems,
	}
}

type DocType interface {
	Transform() interface{}
}

func (toc *TableOfContents) Transform() interface{} {
	return toc.Toc
}

func (nodeDoc *Nodejs) Transform() *Section {
	fmt.Println(len(nodeDoc.DocContent.Nodes))
	nodeDoc.DocContent.Each(func(i int, elem *goquery.Selection) {
		if elem.Nodes[0].Data == "h1" {
			fmt.Println(elem.Text())
		}
	})
	return &Section{}
}
