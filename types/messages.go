package types

import (
	"fmt"
	"strconv"
	"strings"
)

type MessageRules struct {
	rules map[int]MessageRule
}

func NewMessageRules() MessageRules {
	return MessageRules{
		map[int]MessageRule{},
	}
}

func (mr *MessageRules) AddRule(id int, rule MessageRule) {
	mr.rules[id] = rule
}

func (mr *MessageRules) Evaluate(id int) map[string]bool {
	return mr.getEvaluatedMatches(id)
}

func (mr *MessageRules) getEvaluatedMatches(id int) map[string]bool {
	rule := mr.rules[id]
	ruleSet := rule.rule
	return mr.EvaluatRuleSet(ruleSet)
}

func (mr *MessageRules) EvaluatRuleSet(ruleSet []string ) map[string]bool {

	matches := map[string]bool{}

	// base case if we get a rule that has a value
	if ruleSet[0] == "\"" {
		value := ruleSet[1]
		if len(matches) == 0 {
			matches[value] = true
		} else {
			for match,_ := range matches {
				newValue := fmt.Sprintf("%s%s", match, value)
				matches[newValue] = true
			}
		}

		return matches
	}

	hasPipe := false
	for _, ruleIndexString := range ruleSet {
		if ruleIndexString == "|" {
			hasPipe = true
		}
	}

	// basic reference no |
	if !hasPipe {
		for _, ruleIndexString := range ruleSet {
			ruleIndex, err := strconv.Atoi(ruleIndexString)

			// if numeric
			if err == nil {
				newMatches := mr.getEvaluatedMatches(ruleIndex)

				if len(matches) == 0 {
					matches = newMatches
				} else {
					originalMatches := copyMatches(matches)
					for match,_ := range originalMatches {
						for newMatch,_ := range newMatches {
							// Clear old match
							delete(matches, match)
							newValue := fmt.Sprintf("%s%s", match, newMatch)
							matches[newValue] = true
						}
					}
				}
			}
		}
	}

	if hasPipe{
		subRuleSets := strings.Split(strings.Join(ruleSet, " "), "|")

		subMatches := map[string]bool{}
		for _, ruleSet := range subRuleSets {
			newMatches := mr.EvaluatRuleSet(strings.Split(ruleSet, " "))

			for match,_ := range newMatches {
				subMatches[match] = true
			}
		}

		if len(matches) == 0 {
			matches = subMatches
		} else {
			originalMatches := copyMatches(matches)
			for match,_ := range originalMatches {
				for newMatch,_ := range subMatches {
					// Clear old match
					delete(matches, match)
					newValue := fmt.Sprintf("%s%s", match, newMatch)
					matches[newValue] = true
				}
			}
		}
	}

	return matches
}

type MessageRule struct {
	rule []string
}

func NewMessageRule(def []string) MessageRule {
	return MessageRule{
		def,
	}
}

func copyMatches(originalMatches  map[string]bool) map[string]bool {
	newMatches := map[string]bool{}
	for k,v := range  originalMatches {
		newMatches[k] = v
	}

	return newMatches
}
