package parser

import (
	"codoc/types"
	"fmt"
)

type Parsed struct {
	DocInfo   types.Doc
	ParsedDoc string `json:"doc"`
}

func Parse(toParse string, docInfo types.Doc) *Parsed {
	fmt.Println(toParse)
	return &Parsed{
		DocInfo:   docInfo,
		ParsedDoc: "",
	}
}
