package main

import "fmt"

// In Go v1.18 generics were introduced. This program is used to demonstrate an example of generics use-case.
// There are three non-generic functions that add two integers, floats and strings respectively
// and one generic function that serves the same purpose and supports types that satisfy summable interface

func main() {
	var (
		aInt = 5
		bInt = 6

		aFloat = 5.5
		bFloat = 6.6

		aString = "five"
		bString = "six"
	)

	// using three separate functions: addInts(...), addFloats(...) and addStrings(...)
	resInt := addInts(aInt, bInt)
	resFloat := addFloats(aFloat, bFloat)
	resString := addStrings(aString, bString)

	fmt.Printf("\nseparate functions: int sum: %d; float sum: %g; string sum: %s;\n", resInt, resFloat, resString)

	// using one generic function add(...) that supports any type of summable interface
	resInt = add(aInt, bInt)
	resFloat = add(aFloat, bFloat)
	resString = add(aString, bString)

	fmt.Printf("\ngeneric function: int sum: %d; float sum: %g; string sum: %s;\n", resInt, resFloat, resString)
}

// summable is interface that consists of main golang types that support addition.
type summable interface {
	uint | int | float64 | string
}

// add sums a and b of type summable.
func add[T summable](a, b T) T {
	return a + b
}

// addInts sums integers a and b.
func addInts(a, b int) int {
	return a + b
}

// addFloats sums float64s a and b.
func addFloats(a, b float64) float64 {
	return a + b
}

// addStrings concatenates strings a and b.
func addStrings(a, b string) string {
	return a + b
}
