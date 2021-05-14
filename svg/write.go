package svg

import (
	"io"
	"strings"

	"github.com/adnsv/srw/misc/multiline"
	"github.com/adnsv/srw/xml"
)

func Write(out io.Writer, e *Element, cfg xml.WriterConfig) {
	w := xml.NewWriter(out, cfg)
	w.BOM()
	w.XmlDecl()
	WriteElement(w, e)
}

func WriteElement(w *xml.Writer, e *Element) {
	w.OTag(e.Name)
	for _, a := range e.Attrs {
		w.RawAttr(a.Name, a.Value)
	}
	if len(e.Children) > 0 {
		for _, c := range e.Children {
			WriteElement(w, c)
		}
	} else if len(e.Content) > 0 {
		reindent := e.Name.Raw() == "style" || e.Attr("type").Raw() == "text/css"
		if reindent {
			s := e.Content.Unscrambled()
			indent := w.IndentStr()
			s = processContent(s, strings.Repeat(indent, w.IndentLevel()), indent)
			w.String(s)
		} else {
			w.RawString(e.Content)
		}
	}
	w.CTag()
}

func processContent(s string, prefix, indent string) string {
	ss := multiline.SplitAtEOLs(string(s))
	ss = multiline.TrimTopAndBottom(ss)
	multiline.Reindent(ss, prefix+indent, indent)
	return "\n" + strings.Join(ss, "\n") + "\n" + prefix
}
