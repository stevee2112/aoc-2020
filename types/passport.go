package types

import (
)

type Passport struct {
	BirthYear string
	IssueYear string
	ExpirationYear string
	Height string
	HairColor string
	EyeColor string
	PassportId string
	CountryId string
}

func (p *Passport) SetField(field string, value string) {
	switch (field) {
	case "byr":
		p.BirthYear = value
	case "iyr":
		p.IssueYear = value
	case "eyr":
		p.ExpirationYear = value
	case "hgt":
		p.Height = value
	case "hcl":
		p.HairColor = value
	case "ecl":
		p.EyeColor = value
	case "pid":
		p.PassportId = value
	case "cid":
		p.CountryId = value
	}
}

func (p *Passport) IsValid() bool {

	if p.BirthYear == "" {
		return false
	}

	if p.IssueYear == "" {
		return false
	}

	if p.ExpirationYear == "" {
		return false
	}

	if p.Height == "" {
		return false
	}

	if p.HairColor == "" {
		return false
	}

	if p.EyeColor == "" {
		return false
	}

	if p.PassportId == "" {
		return false
	}

	return true
}
