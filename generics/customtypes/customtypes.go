package main

import "fmt"

type celsius float64
type fahrenheit float64

func main() {
	var (
		num      int        = 5
		tempCels celsius    = 232.8
		tempFahr fahrenheit = 451
	)

	// Here if you remove comments before the second and third line in this block you would receive an error.
	// That is because "badNumeric" interface does not support derived types of int: "celsius" and "fahrenheit".
	badPrintWithSymbol(num)
	// badPrintWithSymbol(tempCels)
	// badPrintWithSymbol(tempFahr)

	// However here everything would work nicely because derived types are specified in "goodNumeric" interface
	// using ~ (tilde) symbol.
	goodPrintWithSymbol(num)
	goodPrintWithSymbol(tempCels)
	goodPrintWithSymbol(tempFahr)
}

// badNumeric interface supports only int and float64 types. No derivatives!
type badNumeric interface {
	int | float64
}

// badPrintWithSymbol prints numbers of types specified in badNumeric interface with + or - in front of them.
func badPrintWithSymbol[N badNumeric](n N) {
	switch {
	case n > 0:
		fmt.Printf("\n+%v\n", n)
	case n < 0:
		fmt.Printf("\n-%v\n", n)
	case n == 0:
		fmt.Printf("\n0\n")
	}
}

// goodNumeric interface supports int, float64 and any type derived from them.
// That effect is achieved by using ~ (tilde) symbol.
type goodNumeric interface {
	~int | ~float64
}

// goodPrintWithSymbol prints numbers of types specified in goodNumeric interface with + or - in front of them.
func goodPrintWithSymbol[N goodNumeric](n N) {
	switch {
	case n > 0:
		fmt.Printf("\n+%v\n", n)
	case n < 0:
		fmt.Printf("\n-%v\n", n)
	case n == 0:
		fmt.Printf("\n0\n")
	}
}
