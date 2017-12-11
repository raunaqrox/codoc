package types

type Doc struct {
	DocName string
	DocUrl  string
	DocPath string
}

type DocInputFormat struct {
	Meta   Meta   `json:"meta"`
	ToC    string `json:"table_of_contents"`
	Topic  string `json:"topic"`
	Para   string `json:"para"`
	Format string `json:"format"`
}

type Meta struct {
	url  string
	name string
}

type TocElem struct {
	Name string
	Link string
}

type DocOutputFormat struct {
	TableOfContents TocElem
}

func NewTocElem(name, link string) *TocElem {
	return &TocElem{
		name,
		link,
	}
}
