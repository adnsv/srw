package xml_test

import (
	"bytes"
	"fmt"

	"github.com/adnsv/srw/xml"
)

func Example_write() {

	out := &bytes.Buffer{}
	w := xml.NewWriter(out)

	s := "Hello, "

	w.OTag("hello")
	w.Attr("attr", 42)
	w.Attr("attr2", true)
	w.BeginContent()
	w.Write(s)
	w.Write("World!")
	w.Write(' ')
	w.Write([]string{"a", "b", "c"})
	w.CTag()

	fmt.Println(out.String())
}
