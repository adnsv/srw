package css

import gc "github.com/gorilla/css/scanner"

type Rule struct {
	AtKeywordToken *gc.Token
	Prelude        []ComponentValue
	Block          *Block
}

// Component represents CSS "Component value",
// which is one of PreservedToken, CurlyBlock, RoundBlock, SquareBlock
type ComponentValue interface {
}

type BlockType int

const (
	CurlyBlock = iota
	RoundBlock
	SquareBlock
)

type Block struct {
	Type       BlockType
	Components []ComponentValue // contains Components
}

type Function struct {
	Name       string
	Components []ComponentValue
}

type AtRule struct {
	Identifier string
	Components []ComponentValue
	CurlyBlock *Block
}
