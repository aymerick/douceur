package douceur

import "fmt"

const (
	QUALIFIED_RULE RuleKind = iota
	AT_RULE
)

var atRulesWithRulesBlock = []string{
	"@document", "@font-feature-values", "@keyframes", "@media", "@supports",
}

type RuleKind int

type Rule struct {
	Kind         RuleKind
	Prelude      string
	Declarations []*Declaration

	// At Rule name (eg: "@media")
	Name string

	// At Rule embedded rules
	Rules []*Rule
}

func NewRule(kind RuleKind) *Rule {
	return &Rule{
		Kind: kind,
	}
}

// Returns string representation of rule kind
func (kind RuleKind) String() string {
	switch kind {
	case QUALIFIED_RULE:
		return "Qualified Rule"
	case AT_RULE:
		return "At Rule"
	default:
		return "WAT"
	}
}

// Returns true if this rule embeds another rules
func (rule *Rule) embedsRules() bool {
	if rule.Kind == AT_RULE {
		for _, atRuleName := range atRulesWithRulesBlock {
			if rule.Name == atRuleName {
				return true
			}
		}
	}

	return false
}

func (rule *Rule) String() string {
	result := ""

	// result += fmt.Sprintf("[%s] ", rule.Kind.String())

	if rule.Kind == AT_RULE {
		result += fmt.Sprintf("%s ", rule.Name)
	}

	if rule.Prelude != "" {
		result += fmt.Sprintf("%s", rule.Prelude)
	}

	if (len(rule.Declarations) == 0) && (len(rule.Rules) == 0) {
		result += ";"
	} else {

		result += " {\n"

		if rule.embedsRules() {
			for _, subRule := range rule.Rules {
				result += fmt.Sprintf("  %s\n", subRule.String())
			}
		} else {
			for _, decl := range rule.Declarations {
				result += fmt.Sprintf("  %s\n", decl.String())
			}
		}

		result += "}\n"
	}

	return result
}
