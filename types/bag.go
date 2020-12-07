package types

import (
)

type Rules struct {
	rules map[string]Rule
}

func NewRules() Rules {
	return Rules{
		map[string]Rule{},
	}
}

func (b *Rules) AddRule(rule Rule) {
	b.rules[rule.Color] = rule
}

func (b *Rules) PossibleBags(searchColor string) []Rule {

	rules := []Rule{}
	for color, rule := range b.rules {
		if b.ContainsBagColor(color, searchColor) {
			rules = append(rules, rule)
		}
	}

	return rules
}

func (b *Rules) ContainsBagColor(ruleColor string, searchColor string) bool {

	for childColor, _ := range b.rules[ruleColor].Children {
		if  childColor == searchColor || b.ContainsBagColor(childColor, searchColor) {
			return true
		}
	}

	return false
}


type Rule struct {
	Color string
	Children map[string]int
}

func NewRule(color string, childColors map[string]int) Rule {

	return Rule{color, childColors}
}
