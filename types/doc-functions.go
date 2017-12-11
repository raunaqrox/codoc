package types

func NewTocElem(name, link string) *TocElem {
	return &TocElem{
		name,
		link,
	}
}

func NewDocOutputFormat(todoElems []*TocElem) *DocOutputFormat {
	return &DocOutputFormat{
		todoElems,
	}
}
