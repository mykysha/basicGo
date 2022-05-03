package main

// generic addition introduced interface "comparable" of types that can be compared using == and != operators.
//
// WARNING! Operators > and < are not defined by "comparable" interface!
//
// go.dev page says:
// "comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers,
// channels, arrays of comparable types, structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint, not as the type of a variable"

// RemoveDuplicates drops recurring elements from a slice.
func RemoveDuplicates[T comparable](arr []T) []T {
	var res []T

	res = append(res, arr[0])

	for i := 1; i < len(arr); i++ {
		var duplicate bool

		for _, el := range res {
			if arr[i] == el { // as of GoLand 2021.3.4 this is considered an error, but compiles successfully
				duplicate = true
			}
		}

		if !duplicate {
			res = append(res, arr[i])
		}
	}

	return res
}
