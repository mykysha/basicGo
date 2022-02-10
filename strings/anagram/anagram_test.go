package main

import (
	"fmt"
	"testing"
)

func TestAnagrams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName       string
		firstWord      string
		secondWord     string
		expectedResult bool
	}{
		{"Fist upper palindrome", "Anna", "", true},
		{"All upper palindrome", "RACECAR", "", true},
		{"Upper and lower palindrome", "TeNeT", "", true},
		{"All lower palindrome", "mom", "", true},
		{"Very long palindrome", fmt.Sprint(
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
				"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna"),
			"", true},

		{"First upper not a palindrome", "Word", "", false},
		{"All upper not a palindrome", "CAPS", "", false},
		{"Upper and lower not a palindrome", "UpPeRlOwEr", "", false},
		{"All lower not a palindrome", "alllower", "", false},
		{"Very long not a palindrome", "", fmt.Sprint(
			"Thisphraseisnotsupposedtoberegardedasapalindromecauseitisnotoneasiammakingthisoneup" +
				"sothereisexactlyzerochancethattheresultisgoingtobetruehoweveritisbettertotestthoingsoutjustincase"),
			false},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

		})
	}
}
