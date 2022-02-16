package palindrome_test

import (
	"fmt"
	"testing"

	"github.com/nndergunov/basicGo/strings/palindrome/palindromefinder"
)

func BenchmarkTestSingleThreadPalindrome_regularPalindrome(b *testing.B) {
	regularPalindrome := "racecar"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(regularPalindrome)

		if !res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestSingleThreadPalindrome_regularWord(b *testing.B) {
	regularWord := "raAecar"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(regularWord)

		if res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestSingleThreadPalindrome_longPalindrome(b *testing.B) {
	longPalindrome := fmt.Sprint(
		"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(longPalindrome)

		if !res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestSingleThreadPalindrome_longWord(b *testing.B) {
	longWord := fmt.Sprint(
		"AnnacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatAtenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(longWord)

		if res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestSingleThreadPalindrome_veryLongPalindrome(b *testing.B) {
	veryLongPalindrome := fmt.Sprint(
		"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(veryLongPalindrome)

		if !res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestSingleThreadPalindrome_veryLongWord(b *testing.B) {
	veryLongWord := fmt.Sprint(
		"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"AnnacivickayaklevelmadammomnoonracecarradarredderreferrepaperAotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(veryLongWord)

		if res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestMultiThreadPalindrome_regularPalindrome(b *testing.B) {
	regularPalindrome := "racecar"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(regularPalindrome)

		if !res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestMultiThreadPalindrome_regularWord(b *testing.B) {
	regularWord := "raAecar"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(regularWord)

		if res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestMultiThreadPalindrome_longPalindrome(b *testing.B) {
	longPalindrome := fmt.Sprint(
		"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(longPalindrome)

		if !res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestMultiThreadPalindrome_longWord(b *testing.B) {
	longWord := fmt.Sprint(
		"AnnacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatAtenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(longWord)

		if res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestMultiThreadPalindrome_veryLongPalindrome(b *testing.B) {
	veryLongPalindrome := fmt.Sprint(
		"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(veryLongPalindrome)

		if !res || err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTestMultiThreadPalindrome_veryLongWord(b *testing.B) {
	veryLongWord := fmt.Sprint(
		"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"AnnacivickayaklevelmadammomnoonracecarradarredderreferrepaperAotatorrotorsagassolosstatstenetwow" +
			"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna" +
			"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := palindromefinder.CheckIfPalindrome(veryLongWord)

		if res || err != nil {
			b.Fail()
		}
	}
}

func TestPalindromesSingleThread(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName       string
		word           string
		expectedResult bool
	}{
		{"Fist upper palindrome", "Anna", true},
		{"All upper palindrome", "RACECAR", true},
		{"Upper and lower palindrome", "TeNeT", true},
		{"All lower palindrome", "mom", true},
		{
			"Very long palindrome", fmt.Sprint(
				"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
					"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna"),
			true,
		},

		{"First upper not a palindrome", "Word", false},
		{"All upper not a palindrome", "CAPS", false},
		{"Upper and lower not a palindrome", "UpPeRlOwEr", false},
		{"All lower not a palindrome", "alllower", false},
		{
			"Very long not a palindrome", fmt.Sprint(
				"Thisphraseisnotsupposedtoberegardedasapalindromecauseitisnotoneasiammakingthisoneup" +
					"sothereisexactlyzerochancethattheresultisgoingtobetruehoweveritisbettertotestthoingsoutjustincase"),
			false,
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			res, err := palindromefinder.CheckIfPalindrome(test.word)
			if err != nil {
				t.Errorf("did not expect to receive an error")
			}

			if res != test.expectedResult {
				t.Errorf("expected %v, got %v", test.expectedResult, res)
			}
		})
	}
}

func TestPalindromesMultiThread(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName       string
		word           string
		expectedResult bool
	}{
		{"Fist upper palindrome", "Anna", true},
		{"All upper palindrome", "RACECAR", true},
		{"Upper and lower palindrome", "TeNeT", true},
		{"All lower palindrome", "mom", true},
		{
			"Very long palindrome", fmt.Sprint(
				"Annacivickayaklevelmadammomnoonracecarradarredderreferrepaperrotatorrotorsagassolosstatstenetwow" +
					"wowtenetstatssolossagasrotorrotatorrepaperreferredderradarracecarnoonmommadamlevelkayakcivicanna"),
			true,
		},

		{"First upper not a palindrome", "Word", false},
		{"All upper not a palindrome", "CAPS", false},
		{"Upper and lower not a palindrome", "UpPeRlOwEr", false},
		{"All lower not a palindrome", "alllower", false},
		{
			"Very long not a palindrome", fmt.Sprint(
				"Thisphraseisnotsupposedtoberegardedasapalindromecauseitisnotoneasiammakingthisoneup" +
					"sothereisexactlyzerochancethattheresultisgoingtobetruehoweveritisbettertotestthoingsoutjustincase"),
			false,
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			res, err := palindromefinder.CheckIfPalindromeConcurrently(test.word)
			if err != nil {
				t.Errorf("did not expect to receive an error")
			}

			if res != test.expectedResult {
				t.Errorf("expected %v, got %v", test.expectedResult, res)
			}
		})
	}
}
