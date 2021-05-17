package css

import (
	"fmt"

	gc "github.com/gorilla/css/scanner"
)

func mkerr(msg string, line, col int) error {
	return fmt.Errorf("css [%d,%d]: %s", line, col, msg)
}

type reader struct {
	scn *gc.Scanner
	sot *gc.Token
	cur *gc.Token
}

func openReader(buf string) *reader {
	r := &reader{scn: gc.New(buf)}
	return r
}

func (r *reader) mkerr(msg string) error {
	return mkerr(msg, r.sot.Line, r.sot.Column)
}
func (r *reader) next() bool {
	r.cur = r.scn.Next()
	return !r.done()
}
func (r *reader) charv() string {
	if r.cur.Type == gc.TokenChar {
		return r.cur.Value
	}
	return ""
}
func (r *reader) isSpace() bool {
	return r.cur.Type == gc.TokenS
}
func (r *reader) done() bool {
	return r.cur.Type == gc.TokenS || r.cur.Type == gc.TokenEOF
}
func (r *reader) skipSpace() bool {
	if !r.isSpace() {
		return false
	}
	for r.next() && r.isSpace() {
	}
	return true
}

func (r *reader) run() error {
	for r.next() {
		if r.isSpace() {
			continue
		}
		r.sot= r.cur
		if r.cur.Type == gc.TokenCDO {
			// comment
			for r.next()&&r.cur. != 
			i := findtyp(int(gc.TokenCDC))
			if i < 0 {
				return mkerr("unterminated comment")
			}
		}
	}
}

/*


func (r reader) init(buf string) error {
	scanner := gc.New(buf)
	r.tt = r.tt[:0]
	r.cur, r.sot, r.end = 0, 0, 0

	for {
		t := scanner.Next()
		if t.Type == gc.TokenError {
			return mkerr(t.Value, t.Line, t.Column)
		}
		if t.Type == gc.TokenEOF {
			break
		}
		r.tt = append(r.tt, t)
		fmt.Println(t)
	}

	r.end = len(r.tt)
	return nil
}

func (r *reader) currtk() *gc.Token {
	return r.tt[r.cur]
}

func (r *reader) isSpace(idx int) bool {
	return idx < r.end && r.tt[idx].Type == gc.TokenS
}
func (r *reader) skipSpace() bool {
	if !r.isSpace(r.cur) {
		return false
	}
	r.cur++
	return true
}

func (r *reader) charv(idx int) string {
	if idx < r.end && r.tt[idx].Type == gc.TokenChar {
		return r.tt[idx].Value
	}
	return ""
}

func Read(s string) error {
	r := reader{}
	err := r.init(s)
	if err != nil {
		return err
	}

	for r.cur < r.end {
		if r.skipSpace() {
			continue
		}
		r.sot = r.cur
		t := r.tt[r.cur]
		r.cur++

		if t.Type == gc.TokenCDO {
			// comment
			for r.cur < r.end &&
			i := findtyp(int(gc.TokenCDC))
			if i < 0 {
				return mkerr("unterminated comment")
			}
		}

		if t.Type == gc.TokenAtKeyword {
			// at-keyword-token
			ret := &AtRule{
				Identifier: tt[cur].Value,
			}
			for space(cur) {
				cur++
			}
			c := ""
			for {
				if cur == end {
					return mkerr("missing at-rule terminator")
				}
				c = charv(cur)
				if c == ";" || c == "{" {
					cur++
					break
				}
				ret.Components += tt[cur].Value
				cur++
			}
			if c == ";" {
				cur++
			} else if c != "{" {
				return mkerr("missing at-rule terminator")
			} else {
				for {
					if cur == end {
						return mkerr("missing at-rule {}-block terminator")
					}
					if charv(cur) == "}" {
						cur++
						break
					}
					ret.CurlyBlockContent += tt[cur].Value
					cur++
				}
			}
			continue
		}

		// qualified rule
		if t.Type == gcss.TokenChar {

		}
	}
	return nil
}

*/
