package css

type Rule string

type RuleKind int

const (
	EmptyRule = RuleKind(iota)
	TagRule
	ClassRule
	IdRule
)

func (r Rule) Kind() RuleKind {
	if len(r) == 0 {
		return EmptyRule
	}
	if r[0] == '.' {
		return ClassRule
	}
	if r[0] == '#' {
		return IdRule
	}
	return TagRule
}
