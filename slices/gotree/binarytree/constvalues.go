package binarytree

import "errors"

const (
	boolType    = "bool"
	stringType  = "string"
	intType     = "int"
	runeType    = "rune"
	float64Type = "float64"
)

var (
	ErrNodeExists        = errors.New("node with this value already exists")
	errUnknownType       = errors.New("this tree does not support given type")
	errIncompatibleType  = errors.New("incompatible type")
	errTreeTypeAssertion = errors.New("could not get right type of tree node")
)
