package css

import (
	"fmt"
	"io"

	gc "github.com/gorilla/css/scanner"
)

func Write(w io.Writer, rules []*Rule) {
	for i, rule := range rules {
		if i > 0 {
			fmt.Fprintln(w)
		}
		if rule.AtKeywordToken != nil {
			fmt.Fprint(w, rule.AtKeywordToken.Value)
		}
		writeComponents(w, rule.Prelude)
		if rule.Block == nil {
			fmt.Fprintf(w, ";")
		} else {
			fmt.Fprint(w, "{")
			writeComponents(w, rule.Block.Components)
			fmt.Fprint(w, "}")
		}
	}
}

func writeComponents(w io.Writer, components []ComponentValue) {
	for _, component := range components {
		switch c := component.(type) {
		case *gc.Token:
			fmt.Fprint(w, c.Value)
		case *Block:
			fmt.Fprint(w, c.Type.Prefix())
			writeComponents(w, c.Components)
			fmt.Fprint(w, c.Type.Postfix())
		case *Function:
			fmt.Fprintf(w, "%s(", c.Name)
			writeComponents(w, c.Components)
			fmt.Fprintf(w, ")")
		}
	}
}

func (bt BlockType) Prefix() string {
	switch bt {
	case CurlyBlock:
		return "{"
	case RoundBlock:
		return "("
	case SquareBlock:
		return "["
	default:
		return "??"
	}
}
func (bt BlockType) Postfix() string {
	switch bt {
	case CurlyBlock:
		return "}"
	case RoundBlock:
		return ")"
	case SquareBlock:
		return "]"
	default:
		return "??"
	}
}
