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
	Name        string // name of toc element
	Link        string // href of associated anchor tag
	TocPageData NullDocPage
}

// indicating that the page would be null when valid is false
type NullDocPage struct {
	LocalToc *TableOfContents
	MetaData interface{}
	Topic    Section
	Valid    bool // added like NullString, to know if the value is present at the moment or not
}

type Explanation struct {
	Explanation string
	Example     string
}

type Section struct {
	Topic       string
	Explanation []Explanation
}

type TableOfContents struct {
	Toc []*TocElem `json:"Toc"`
}

type DocOutputFormat struct {
	Toc *TableOfContents
	// add [] Section
}

type Parsed struct {
	DocInfo   Doc
	ParsedDoc *DocOutputFormat
}
