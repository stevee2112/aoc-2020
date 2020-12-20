package types

import (
)

type TicketFieldRule struct {
	Name string
	Range1 FieldRange
	Range2 FieldRange
}

func NewFieldRule(name string, min1 int, max1 int, min2 int, max2 int) TicketFieldRule{
	return TicketFieldRule{
		name,
		FieldRange{min1, max1},
		FieldRange{min2, max2},
	}
}

func (t *TicketFieldRule) Match(ticketField TicketField) bool {
	value := ticketField.Value

	if (value >= t.Range1.Min && value <= t.Range1.Max ||
		value >= t.Range2.Min && value <= t.Range2.Max) {
		return true
	} else {
		return false
	}
}

type FieldRange struct {
	Min int
	Max int
}

type TicketRules struct {
	Fields []TicketFieldRule
}

type TicketField struct {
	Value int
	RulesMatched []TicketFieldRule
}

type Ticket struct {
	Fields []TicketField
}

func NewTicket(values []int) Ticket {

	ticket := Ticket{}

	for _,value := range values {
		ticket.Fields = append(ticket.Fields, TicketField{value, []TicketFieldRule{}})
	}

	return ticket
}

func (t *Ticket) MatchFieldsToRules(rules TicketRules) {

	for index,field := range t.Fields{ // Ticket fields
		for _, rule := range rules.Fields { // Field Rules
			if rule.Match(field) {
				t.Fields[index].RulesMatched = append(t.Fields[index].RulesMatched, rule)
			}
		}
	}
}
