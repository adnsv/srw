package css

type Rule struct {
	AtKeywordToken *
}

// Component represents CSS "Component value",
// which is one of PreservedToken, CurlyBlock, RoundBlock, SquareBlock
type ComponentValue interface {
}

type PreservedToken struct {
	ComponentValue
	Components []ComponentValue
}

type BlockType int

type BlockBase struct {
	Components []ComponentValue
}

const (
	CurlyBlock = iota
	RoundBlock
	SquareBlock
)

type Block struct {
	Type           BlockType
	BlockBase      // contains Components
	ComponentValue // can be a child in ComponentValue
}

type FunctionBlock struct {
	//
	Components []ComponentValue
	ComponentValue
}

type AtRule struct {
	Identifier string
	Components []ComponentValue
	CurlyBlock *BlockBase
}
