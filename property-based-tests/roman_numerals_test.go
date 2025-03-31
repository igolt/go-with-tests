package roman_numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var testCases = []struct {
	arabic uint16
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{13, "XIII"},
	{14, "XIV"},
	{15, "XV"},
	{17, "XVII"},
	{19, "XIX"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{88, "LXXXVIII"},
	{90, "XC"},
	{100, "C"},
	{159, "CLIX"},
	{193, "CXCIII"},
	{400, "CD"},
	{500, "D"},
	{798, "DCCXCVIII"},
	{900, "CM"},
	{1000, "M"},
	{1006, "MVI"},
	{1984, "MCMLXXXIV"},
	{2014, "MMXIV"},
	{3999, "MMMCMXCIX"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range testCases {
		testName := fmt.Sprintf("%d gets converted to %s", test.arabic, test.roman)
		t.Run(testName, func(t *testing.T) {
			got := ToRomanNumerals(test.arabic)
			if got != test.roman {
				t.Errorf("expected %q but got %q", test.roman, got)
			}
		})
	}
}

func TestFromRomanNumerals(t *testing.T) {
	for _, test := range testCases {
		testName := fmt.Sprintf("%s gets converted to %d", test.roman, test.arabic)
		t.Run(testName, func(t *testing.T) {
			got := FromRomanNumerals(test.roman)
			if got != test.arabic {
				t.Errorf("expected %d but got %d", test.arabic, got)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)

		roman := ToRomanNumerals(arabic)
		fromRoman := FromRomanNumerals(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
