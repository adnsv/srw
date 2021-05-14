package svg

import (
	"fmt"
	"strings"

	"github.com/adnsv/srw/xml"
)

func FromString(buf string) (*Element, error) {
	cc := xml.Open(buf)
	if !cc.NextTag() {
		return nil, cc.Err()
	}
	if cc.Name() != "svg" {
		return nil, cc.MakeError("svg", fmt.Sprintf("unsupported root element '%s'", cc.Name()))
	}
	ret := &Element{Name: "svg"}
	cc.HandleTag(func(attrs xml.AttributeList, content *xml.Content) error {
		return read(ret, attrs, content)
	})
	if cc.Err() != nil {
		return nil, cc.Err()
	}
	return ret, nil
}

func read(e *Element, aa xml.AttributeList, cc *xml.Content) error {
	e.Attrs = make([]Attr, len(aa))
	for i, a := range aa {
		e.Attrs[i].Name = a.Name
		e.Attrs[i].Value = a.Value
	}
	if cc != nil {
		var contentStr xml.RawString
		collectContentStr := true
		for cc.Next() {
			switch cc.Kind() {
			case xml.Tag:
				collectContentStr = false
				child := &Element{Name: cc.Name()}
				cc.HandleTag(func(attrs xml.AttributeList, content *xml.Content) error {
					return read(child, attrs, content)
				})
				if cc.Err() != nil {
					return cc.Err()
				}
				e.Children = append(e.Children, child)
			case xml.SData:
				if collectContentStr {
					contentStr += cc.Value()
				}
			case xml.CData:
				collectContentStr = false
			case xml.Comment:
				collectContentStr = false
			case xml.PI:
				collectContentStr = false
			}
		}
		if collectContentStr {
			if e.Name != "text" {
				if strings.TrimSpace(string(contentStr)) == "" {
					contentStr = ""
				}
			}
			e.Content = contentStr
		}
	}
	return nil
}
