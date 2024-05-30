package main

import (
	"fmt"
	"strings"
)

const MaxAllowedArabic uint16 = 3999

type romanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(n uint16) (string, error) {
	if n > MaxAllowedArabic {
		return "", fmt.Errorf("max allowed arabic value is %d", MaxAllowedArabic)
	}

	var roman strings.Builder

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return roman.String(), nil
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
