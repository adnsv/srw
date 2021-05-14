package css

import (
	"fmt"

	gcss "github.com/gorilla/css/scanner"
)

func Read(s string) error {
	scanner := gcss.New(s)
	tt := []*gcss.Token{}

	mkerrbase := func(msg string, line, col int) error {
		return fmt.Errorf("css [%d,%d]: %s", line, col, msg)
	}

	for {
		t := scanner.Next()
		if t.Type == gcss.TokenError {
			return mkerrbase(t.Value, t.Line, t.Column)
		}
		if t.Type == gcss.TokenEOF {
			break
		}
		tt = append(tt, t)
	}

	cur, sot, end := 0, 0, len(tt)

	mkerr := func(msg string) error {
		return mkerrbase(msg, tt[sot].Line, tt[sot].Column)
	}

	find := func(typ int) int {
		i := cur
		for ; i < end; i++ {
			if int(tt[i].Type) == typ {
				return i
			}
		}
		return -1
	}

	for cur < end {
		sot = cur
		t := tt[cur]
		cur++
		if t.Type == gcss.TokenS {
			continue
		}

		if t.Type == gcss.TokenCDO {
			// comment
			i := find(int(gcss.TokenCDC))
			if i < 0 {
				return mkerr("unterminated comment")
			}
		}

	}
	return nil
}
