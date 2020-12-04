package types

import (
	"strconv"
	"regexp"
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

func (p *Passport) HasRequiredFields() bool {
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

func (p *Passport) IsValid() bool {

	// four digits; at least 1920 and at most 2002.
	birthYearInt, _ := strconv.Atoi(p.BirthYear)
	if len(p.BirthYear) < 4 || birthYearInt < 1920 || birthYearInt > 2002 {
		return false
	}

	// four digits; at least 2010 and at most 2020.
	issueYearInt, _ := strconv.Atoi(p.IssueYear)
	if len(p.IssueYear) < 4 || issueYearInt < 2010 || issueYearInt > 2020 {
		return false
	}

	// four digits; at least 2020 and at most 2030.
	expirationYearInt, _ := strconv.Atoi(p.ExpirationYear)
	if len(p.ExpirationYear) < 4 || expirationYearInt < 2020 || expirationYearInt > 2030 {
		return false
	}

	// a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	heightType := reg.ReplaceAllString(p.Height, "")

	reg, _ = regexp.Compile("[^0-9]+")
	heightInt, _ := strconv.Atoi(reg.ReplaceAllString(p.Height, ""))
	if heightType != "cm" && heightType != "in" {
		return false
	}

	if heightType == "cm" && (heightInt < 150 || heightInt > 193) {
		return false
	}

	if heightType == "in" && (heightInt < 59 || heightInt > 76) {
		return false
	}

	// a # followed by exactly six characters 0-9 or a-f.
	reg, _ = regexp.Compile("^#([a-fA-F0-9]{6})")
	if !reg.MatchString(p.HairColor) {
		return false
	}

	//exactly one of: amb blu brn gry grn hzl oth.
	if p.EyeColor != "amb" &&
		p.EyeColor != "blu" &&
		p.EyeColor != "brn" &&
		p.EyeColor != "gry" &&
		p.EyeColor != "grn" &&
		p.EyeColor != "hzl" &&
		p.EyeColor != "oth" {
		return false
	}

	// a nine-digit number, including leading zeroes.
	reg, _ = regexp.Compile("[^0-9]+")
	passportId  := reg.ReplaceAllString(p.PassportId, "")
	if len(passportId) != 9 {
		return false
	}

	return true
}
