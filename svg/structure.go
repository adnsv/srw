package svg

import "github.com/adnsv/srw/xml"

type Element struct {
	Name     xml.NameString
	Attrs    []Attr
	Children []*Element
	Content  xml.RawString
}

type Attr struct {
	Name  xml.NameString
	Value xml.RawString
}

func (e *Element) Attr(name xml.NameString) xml.RawString {
	for _, a := range e.Attrs {
		if a.Name == name {
			return a.Value
		}
	}
	return ""
}
