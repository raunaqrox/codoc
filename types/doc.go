package types

type Doc struct {
	DocName string
	DocUrl  string
	DocPath string
}

type DocInputFormat struct {
	Meta   Meta   `json:"meta"`
	Toc    string `json:"table_of_contents"`
	Topic  string `json:"topic"`
	Para   string `json:"para"`
	Format string `json:"format"`
}

type Meta struct {
	url  string
	name string
}

// Each element of table of content
type TocElem struct {
	Name string // name of toc element
	Link string // href of associated anchor tag
	// maybe even the topic to which it links
}

type DocOutputFormat struct {
	TableOfContents []*TocElem
}
