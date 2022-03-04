package binarytree

import (
	"fmt"
	"strings"
)

func (b *BinaryNode) insertBool(val bool) error {
	if b.treeType != boolType {
		if b.treeType != "" {
			return fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, boolType)
		}

		b.value = val
		b.treeType = boolType

		return nil
	}

	nodeVal, ok := b.value.(bool)
	if !ok {
		return errTreeTypeAssertion
	}

	if val == nodeVal {
		return ErrNodeExists
	}

	if val {
		if b.rightNode == nil {
			b.rightNode = newBinaryNode(val, boolType)

			return nil
		}

		return b.rightNode.Insert(val)
	}

	if b.leftNode == nil {
		b.leftNode = newBinaryNode(val, boolType)

		return nil
	}

	return b.leftNode.Insert(val)
}

func (b *BinaryNode) insertString(val string) error {
	val = strings.ToLower(val)

	if b.treeType != stringType {
		if b.treeType != "" {
			return fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, stringType)
		}

		b.value = val
		b.treeType = stringType

		return nil
	}

	nodeVal, ok := b.value.(string)
	if !ok {
		return errTreeTypeAssertion
	}

	comp := strings.Compare(nodeVal, val)

	if comp == 0 {
		return ErrNodeExists
	}

	if comp == -1 {
		if b.rightNode == nil {
			b.rightNode = newBinaryNode(val, stringType)

			return nil
		}

		return b.rightNode.Insert(val)
	}

	if comp == 1 {
		if b.leftNode == nil {
			b.leftNode = newBinaryNode(val, stringType)

			return nil
		}

		return b.leftNode.Insert(val)
	}

	return nil
}

func (b *BinaryNode) insertInt(val int) error {
	if b.treeType != intType {
		if b.treeType != "" {
			return fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, intType)
		}

		b.value = val
		b.treeType = intType

		return nil
	}

	nodeVal, ok := b.value.(int)
	if !ok {
		return errTreeTypeAssertion
	}

	if val == nodeVal {
		return ErrNodeExists
	}

	if val > nodeVal {
		if b.rightNode == nil {
			b.rightNode = newBinaryNode(val, intType)

			return nil
		}

		return b.rightNode.Insert(val)
	}

	if b.leftNode == nil {
		b.leftNode = newBinaryNode(val, intType)

		return nil
	}

	return b.leftNode.Insert(val)
}

func (b *BinaryNode) insertRune(val rune) error {
	if b.treeType != runeType {
		if b.treeType != "" {
			return fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, runeType)
		}

		b.value = val
		b.treeType = runeType

		return nil
	}

	nodeVal, ok := b.value.(rune)
	if !ok {
		return errTreeTypeAssertion
	}

	if val == nodeVal {
		return ErrNodeExists
	}

	if val > nodeVal {
		if b.rightNode == nil {
			b.rightNode = newBinaryNode(val, runeType)

			return nil
		}

		return b.rightNode.Insert(val)
	}

	if b.leftNode == nil {
		b.leftNode = newBinaryNode(val, runeType)

		return nil
	}

	return b.leftNode.Insert(val)
}

func (b *BinaryNode) insertFloat64(val float64) error {
	if b.treeType != float64Type {
		if b.treeType != "" {
			return fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, float64Type)
		}

		b.value = val
		b.treeType = float64Type

		return nil
	}

	nodeVal, ok := b.value.(float64)
	if !ok {
		return errTreeTypeAssertion
	}

	if val == nodeVal {
		return ErrNodeExists
	}

	if val > nodeVal {
		if b.rightNode == nil {
			b.rightNode = newBinaryNode(val, float64Type)

			return nil
		}

		return b.rightNode.Insert(val)
	}

	if b.leftNode == nil {
		b.leftNode = newBinaryNode(val, float64Type)

		return nil
	}

	return b.leftNode.Insert(val)
}
