package main

import "fmt"

type numeric interface {
	int | uint | float64
}

type textish interface {
	string | []rune | []byte
}

// mathStruct is a struct that uses one generic type.
type mathStruct[T numeric] struct {
	firstVal  T
	secondVal T
}

// doMath prints results of 4 basic mathematical operations performed on struct fields.
// Please note [T] after struct declaration: it is essential, but can have any name.
func (m mathStruct[T]) doMath() {
	var bigger, smaller T

	if m.firstVal > m.secondVal { // as of GoLand 2021.3.4 this is considered an error, but compiles successfully
		bigger = m.firstVal
		smaller = m.secondVal
	} else {
		bigger = m.secondVal
		smaller = m.firstVal
	}

	fmt.Printf("\n")
	fmt.Printf("sum: %v\n", bigger+smaller)
	fmt.Printf("difference: %v\n", bigger-smaller)
	fmt.Printf("product: %v\n", bigger*smaller)
	fmt.Printf("quotient: %v\n", bigger/smaller)
}

type human[N numeric, T textish] struct {
	age  N
	name T
}

// say is a method of a human struct.
// Please note: methods in Go do not support generic types that are not used in struct itself
// (cannot declare another [T customType]).
func (h human[N, T]) say(phrase T) {
	fmt.Printf("\n%s, %v says '%s'\n", h.name, h.age, phrase)
}

func main() {
	// single typed generic struct example
	m := mathStruct[int]{
		firstVal:  1,
		secondVal: 2,
	}

	m.doMath()

	// multiple typed generic struct example
	h := human[int, string]{
		age:  18,
		name: "Mykyta",
	}

	h.say("yay, generics!")

	// array of structs with generic type field
	// as it seems, it is not possible to create an array with generics of multiple types
	mathArr := []mathStruct[int]{
		{
			firstVal:  10,
			secondVal: 2,
		},
		{
			firstVal:  7,
			secondVal: 14,
		},
	}

	for _, s := range mathArr {
		s.doMath()
	}

	// however, there is a valid workaround using our old friend empty interface
	multipleTArr := []interface{}{
		mathStruct[int]{
			firstVal:  55,
			secondVal: 11,
		},
		mathStruct[float64]{
			firstVal:  5.5,
			secondVal: 1.1,
		},
	}

	for _, el := range multipleTArr {
		// of course, we have to use type switch
		switch mathS := el.(type) {
		case mathStruct[int]:
			mathS.doMath()
		case mathStruct[float64]:
			mathS.doMath()
		case mathStruct[uint]:
			mathS.doMath()
		}
	}
}
