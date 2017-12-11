package types

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

type Handlers interface {
	Transform() interface{}
}

func (toc *TableOfContents) Transform() interface{} {
	return toc.Toc
}
