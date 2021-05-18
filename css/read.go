package css

import (
	"fmt"

	gc "github.com/gorilla/css/scanner"
)

func Read(buf string) ([]*Rule, error) {
	r := openReader(buf)
	return r.consumeListOfRules()
}

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
func (r *reader) eof() bool {
	return r.cur.Type == gc.TokenEOF
}
func (r *reader) done() bool {
	return r.cur.Type == gc.TokenError || r.cur.Type == gc.TokenEOF
}
func (r *reader) skipSpace() bool {
	if !r.isSpace() {
		return false
	}
	for r.next() && r.isSpace() {
	}
	return true
}

func (r *reader) consumeListOfRules() ([]*Rule, error) {
	rules := []*Rule{}
	for r.next() {
		r.sot = r.cur

		switch r.cur.Type {

		case gc.TokenS:
			continue

		case gc.TokenEOF:
			return rules, nil

		case gc.TokenCDO:
			for r.next() {
				switch r.cur.Type {
				case gc.TokenCDC:
					continue
				case gc.TokenEOF:
					return nil, r.mkerr("unterminated comment")
				}
			}
		case gc.TokenAtKeyword:
			rule, err := r.consumeRule(true)
			if err != nil {
				return nil, err
			}
			rules = append(rules, rule)

		default:
			rule, err := r.consumeRule(false)
			if err != nil {
				return nil, err
			}
			rules = append(rules, rule)
		}
	}
	return rules, nil
}

func (r *reader) consumeRule(atrule bool) (*Rule, error) {
	rule := &Rule{}
	if atrule {
		rule.AtKeywordToken = r.cur
		if !r.next() {
			return nil, r.mkerr("unterminated at-rule")
		}
	}
	for ; !r.done(); r.next() {
		switch {
		case r.charv() == ";":
			r.next()
			return rule, nil
		case r.cur.Type == gc.TokenEOF:
			return rule, r.mkerr("unterminated at-rule")
		case r.charv() == "{":
			blk, err := r.consumeSimpleBlock()
			if err != nil {
				return nil, err
			}
			rule.Block = blk
			return rule, nil
		default:
			cv, err := r.consumeComponentValue()
			if err != nil {
				return nil, err
			}
			rule.Prelude = append(rule.Prelude, cv)
		}
	}
	return rule, nil
}

func (r *reader) consumeSimpleBlock() (*Block, error) {
	blk := &Block{}
	switch r.charv() {
	case "{":
		blk.Type = CurlyBlock
	case "(":
		blk.Type = RoundBlock
	case "[":
		blk.Type = SquareBlock
	default:
		return nil, r.mkerr("invalid block token")
	}
	term := blk.Type.Postfix()
	for r.next() {
		switch {
		case r.charv() == term:
			return blk, nil
		case r.eof():
			return nil, r.mkerr("unterminated " + term + "-block")
		default:
			c, err := r.consumeComponentValue()
			if err != nil {
				return nil, err
			}
			blk.Components = append(blk.Components, c)
		}
	}
	return blk, nil
}

func (r *reader) consumeComponentValue() (ComponentValue, error) {
	c := r.charv()
	switch {
	case c == "{", c == "(", c == "[":
		return r.consumeSimpleBlock()
	case r.cur.Type == gc.TokenFunction:
		return r.consumeFunction()
	default:
		return r.cur, nil
	}
}

func (r *reader) consumeFunction() (*Function, error) {
	f := &Function{Name: r.cur.Value}
	for r.next() {
		switch {
		case r.charv() == ")":
			return f, nil
		case r.eof():
			return nil, r.mkerr("unterminated function")
		default:
			c, err := r.consumeComponentValue()
			if err != nil {
				return nil, err
			}
			f.Components = append(f.Components, c)
		}
	}
	return f, nil
}
