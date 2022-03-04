package anagram_test

import (
	"fmt"
	"testing"

	"github.com/nndergunov/basicGo/strings/anagram/anagramfinder"
)

func TestAnagrams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName       string
		firstWord      string
		secondWord     string
		expectedResult bool
	}{
		{"Fist pair of anagrams", "meals", "males", true},
		{"Second pair of anagrams", "saint", "satin", true},
		{"Third pair of anagrams", "avenge", "Geneva", true},
		{"Fourth pair of anagrams", "meals", "Salem", true},
		{
			"Long anagrams",
			fmt.Sprint(
				"mealssaintavengemealssalesbalmmeansaltsblotmelonblow" +
					"moistsharpbragmoreshrubchumneedssirencoalnervedskidscounts" +
					"noneskillnudesnaildiaryoceansoberdomainspacesoildottierpairs" +
					"solofiredpalesprayfringepanelsstackhastenparksstickicedpools" +
					"stripinchportsstudykeenpoststeamlampracestooledlastreapvoteslimped" +
					"reefwaitslionrobedwaspslootedrockwellslumproomwestmarchropeswhatmash"),
			fmt.Sprint(
				"malessatinGenevaSalemsealslambmanelastsboltlemonbowl" +
					"omitsharpsgrabRomebrushmuchdenserinsecolaDenverdisksTucson" +
					"neonkillsdunenailsdairycanoerobesMadisoncapeoilsDetroitParis" +
					"OslofriedleappraysfingerNaplestacksAthenssparkticksdicespool" +
					"tripschinsportdustykneestopsmeatpalmcaresToledosaltpearstovedimple" +
					"freewaistloinboredswapsToledocorkswellplummoorstewcharmporethawshams"),
			true,
		},

		{"Not anagrams with different length", "word", "spaceship", false},
		{"Not anagrams with same length", "music", "panda", false},
		{"Not anagrams with same letters visually in eng and rus", "apec", "арес", false},
		{"Not anagrams with same byte length", "pillow", "три", false},
		{
			"Long almost anagrams",
			fmt.Sprint(
				"mealssaintavengemealssalesbalmmeansaltsblotmelonblow" +
					"moistsharpbragmoreshrubchumneedssirencoalnervedskidscounts" +
					"noneskillnudesnaildiaryoceansoberdomainspacesoildottierpairs" +
					"solofiredpalesprayfringepanelsstackhastenparksstickicedpools" +
					"stripinchportsstudykeenpoststeamlampracestooledlastreapvoteslimped" +
					"reefwaitslionrobedwaspslootedrockwellslumproomwestmarchropeswhatmash"),
			fmt.Sprint(
				"malessatinGenevaSalemsealslambmanelastsboltlemonbowl" +
					"omitsharpsgrabRomebrushmuchdenserinsecolaDenverdisksTucson" +
					"neonkillsdunenailsdairycanoerobesMadisoncapeoilsDetroitParis" +
					"OslofriedleappraysfingerNaplestacksAthenssparkticksdicespool" +
					"tripschinsportdustykneestopsmeatpalmcaresToledosaltpearstovedimple" +
					"freewaistloinboredswapsToledocorkswellplummoorstewcharmporethawshame"),
			false,
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			res, err := anagramfinder.CheckIfAnagram(test.firstWord, test.secondWord)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if res != test.expectedResult {
				t.Errorf("ecpected %v, got %v", test.expectedResult, res)
			}
		})
	}
}
