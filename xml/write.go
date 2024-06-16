package xml

import (
	"io"
	"reflect"
	"strings"
)

type Indent int

const (
	IndentTabs    = Indent(0)
	Indent2Spaces = Indent(2)
	Indent4Spaces = Indent(4)
	IndentNone    = Indent(-1)
)

type WriterConfig struct {
	Indent    Indent
	IndentAll bool // force indent all tags
}

type Writer struct {
	WriterConfig
	out           io.Writer
	names         []NameString
	inOTag        bool
	indentLevel   int
	prevLineLevel int
}

func NewWriter(out io.Writer, config WriterConfig) *Writer {
	return &Writer{WriterConfig: config, out: out}
}

func (w *Writer) IndentLevel() int {
	return w.indentLevel
}

func (w *Writer) IndentStr() string {
	switch w.Indent {
	case IndentTabs:
		return "\t"
	case Indent2Spaces:
		return "  "
	case Indent4Spaces:
		return "    "
	default:
		if w.Indent > 0 {
			return strings.Repeat(" ", int(w.Indent))
		}
		return ""
	}
}

func (w *Writer) put(s RawString) {
	w.out.Write([]byte(s))
}

func (w *Writer) BeginContent() {
	if w.inOTag {
		w.put(">")
		w.inOTag = false
	}
}

func (w *Writer) putIndent(level int) {
	if w.Indent == IndentTabs {
		writeTabs(w.out, level)
	} else if w.Indent > 0 {
		writeSpaces(w.out, level*int(w.Indent))
	}
}

func (w *Writer) BOM() {
	w.put("\xef\xbb\xbf")
}

func (w *Writer) XmlDecl() {
	if len(w.names) > 0 {
		panic("xml writer: invalid XmlDecl placement")
	}
	w.put("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
}

func (w *Writer) XmlStandaloneDecl() {
	if len(w.names) > 0 {
		panic("xml writer: invalid XmlDecl placement")
	}
	w.put("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n")
}

const nolevel = -1

func (w *Writer) OTag(name NameString) *Writer {
	if len(name) == 0 || name == "+" {
		panic("xml writer: trying to write a tag with empty name")
	}
	prevLevel := w.prevLineLevel
	w.BeginContent()
	indent := name[0] == '+'
	w.names = append(w.names, name)
	if indent {
		name = name[1:]
	}
	if indent || w.IndentAll {
		if len(w.names) <= 1 {
			w.indentLevel = 0
		} else {
			w.indentLevel++
		}
		if prevLevel == nolevel || prevLevel > w.indentLevel {
			w.put("\n")
			w.putIndent(w.indentLevel)
		} else {
			w.putIndent(w.indentLevel - prevLevel)
		}
	}
	w.put("<")
	w.put(RawString(name.Raw()))
	w.inOTag = true
	w.prevLineLevel = nolevel
	return w
}

func (w *Writer) CTag() *Writer {
	if len(w.names) == 0 {
		panic("xml writer: tag stack underflow")
	}
	name := w.names[len(w.names)-1]
	w.names = w.names[:len(w.names)-1]

	indented := name[0] == '+'
	if indented {
		name = name[1:]
	}
	if w.inOTag {
		w.inOTag = false
		w.put("/>")
	} else {
		w.put("</")
		w.put(RawString(name.Raw()))
		w.put(">")
	}
	if indented || w.IndentAll {
		if w.indentLevel > 0 {
			w.indentLevel--
		}
		w.put("\n")
		w.putIndent(w.indentLevel)
		w.prevLineLevel = w.indentLevel
	}
	return w
}

func (w *Writer) String(s string) *Writer {
	w.BeginContent()
	w.put(ScrambleContent(s))
	return w
}

func (w *Writer) RawString(s RawString) *Writer {
	w.BeginContent()
	w.put(s)
	return w
}

func (w *Writer) Write(v interface{}) *Writer {
	w.BeginContent()
	toContent(w, v)
	return w
}

func (w *Writer) Comment(s string) *Writer {
	w.BeginContent()
	w.put("<!--")
	w.put(RawString(strings.ReplaceAll(s, "--", "-"))) // make sure double-dash is not written
	w.put("-->")
	return w
}

func (w *Writer) StringAttr(name NameString, value string) *Writer {
	if !w.inOTag {
		panic("xml writer: trying to write an attribute outside of an open tag")
	}
	w.put(" ")
	w.put(RawString(name.Raw()))
	w.put(`="`)
	w.put(ScrambleAttr(value))
	w.put(`"`)
	return w
}

func (w *Writer) RawAttr(name NameString, value RawString) *Writer {
	if !w.inOTag {
		panic("xml writer: trying to write an attribute outside of an open tag")
	}
	w.put(" ")
	w.put(RawString(name.Raw()))
	w.put(`="`)
	w.put(value)
	w.put(`"`)
	return w
}

func (w *Writer) Attr(name NameString, value interface{}) *Writer {
	if !w.inOTag {
		panic("xml writer: trying to write an attribute outside of an open tag")
	}
	s, _ := toRawStr(ScrambleAttr, value)
	w.put(" ")
	w.put(RawString(name.Raw()))
	w.put(`="`)
	w.put(s)
	w.put(`"`)
	return w
}

func (w *Writer) OptStringAttr(name NameString, value string) *Writer {
	if !w.inOTag {
		panic("xml writer: trying to write an attribute outside of an open tag")
	}
	if len(value) == 0 {
		return w
	}
	w.put(" ")
	w.put(RawString(name.Raw()))
	w.put(`="`)
	w.put(ScrambleAttr(value))
	w.put(`"`)
	return w
}

func (w *Writer) OptRawAttr(name NameString, value RawString) *Writer {
	if !w.inOTag {
		panic("xml writer: trying to write an attribute outside of an open tag")
	}
	if len(value) == 0 {
		return w
	}
	w.put(" ")
	w.put(RawString(name.Raw()))
	w.put(`="`)
	w.put(value)
	w.put(`"`)
	return w
}

func (w *Writer) OptAttr(name NameString, value interface{}) *Writer {
	if !w.inOTag {
		panic("xml writer: trying to write an attribute outside of an open tag")
	}
	r, _ := toRawStr(ScrambleAttr, value)
	if len(r) == 0 {
		return w
	}
	w.put(" ")
	w.put(RawString(name.Raw()))
	w.put(`="`)
	w.put(r)
	w.put(`"`)
	return w
}

var tabs = [8]byte{'\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t'}

const spaces = "                                "

func writeTabs(w io.Writer, n int) (err error) {
	bb := tabs[:]
	for n > 8 {
		_, err = w.Write(bb)
		if err != nil {
			return
		}
		n -= 8
	}
	if n > 0 {
		_, err = w.Write(bb[:n])
	}
	return
}

func writeSpaces(w io.Writer, n int) (err error) {
	ns := len(spaces)
	bb := []byte(spaces)
	for n > ns {
		_, err = w.Write(bb)
		if err != nil {
			return
		}
		n -= ns
	}
	if n > 0 {
		_, err = w.Write(bb[:n])
	}
	return
}

// UnsupportedTypeError is returned when Marshal encounters a type
// that cannot be converted into XML.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "xml: unsupported type: " + e.Type.String()
}
