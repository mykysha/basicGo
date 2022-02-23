package binarytree

import (
	"fmt"
	"strings"
)

func (b *BinaryNode) boolExists(val bool) (bool, error) {
	if b.treeType != boolType {
		if b.treeType != "" {
			return false, fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, boolType)
		}

		return false, nil
	}

	nodeVal, ok := b.value.(bool)
	if !ok {
		return false, errTreeTypeAssertion
	}

	if val == nodeVal {
		return true, nil
	}

	if val {
		if b.rightNode == nil {
			return false, nil
		}

		return b.rightNode.ValueExists(val)
	}

	if b.leftNode == nil {
		return false, nil
	}

	return b.leftNode.ValueExists(val)
}

func (b *BinaryNode) stringExists(val string) (bool, error) {
	if b.treeType != stringType {
		if b.treeType != "" {
			return false, fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, stringType)
		}

		return false, nil
	}

	nodeVal, ok := b.value.(string)
	if !ok {
		return false, errTreeTypeAssertion
	}

	val = strings.ToLower(val)

	comp := strings.Compare(nodeVal, val)

	if comp == 0 {
		return true, nil
	}

	if comp == -1 {
		if b.rightNode == nil {
			return false, nil
		}

		return b.rightNode.ValueExists(val)
	}

	if comp == 1 {
		if b.leftNode == nil {
			return false, nil
		}

		return b.leftNode.ValueExists(val)
	}

	return false, nil
}

func (b *BinaryNode) intExists(val int) (bool, error) {
	if b.treeType != intType {
		if b.treeType != "" {
			return false, fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, intType)
		}

		return false, nil
	}

	nodeVal, ok := b.value.(int)
	if !ok {
		return false, errTreeTypeAssertion
	}

	if val == nodeVal {
		return true, nil
	}

	if val > nodeVal {
		if b.rightNode == nil {
			return false, nil
		}

		return b.rightNode.ValueExists(val)
	}

	if b.leftNode == nil {
		return false, nil
	}

	return b.leftNode.ValueExists(val)
}

func (b *BinaryNode) runeExists(val rune) (bool, error) {
	if b.treeType != runeType {
		if b.treeType != "" {
			return false, fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, runeType)
		}

		return false, nil
	}

	nodeVal, ok := b.value.(rune)
	if !ok {
		return false, errTreeTypeAssertion
	}

	if val == nodeVal {
		return true, nil
	}

	if val > nodeVal {
		if b.rightNode == nil {
			return false, nil
		}

		return b.rightNode.ValueExists(val)
	}

	if b.leftNode == nil {
		return false, nil
	}

	return b.leftNode.ValueExists(val)
}

func (b *BinaryNode) float64Exists(val float64) (bool, error) {
	if b.treeType != float64Type {
		if b.treeType != "" {
			return false, fmt.Errorf("%w. should be %s, got %s", errIncompatibleType, b.treeType, float64Type)
		}

		return false, nil
	}

	nodeVal, ok := b.value.(float64)
	if !ok {
		return false, errTreeTypeAssertion
	}

	if val == nodeVal {
		return true, nil
	}

	if val > nodeVal {
		if b.rightNode == nil {
			return false, nil
		}

		return b.rightNode.ValueExists(val)
	}

	if b.leftNode == nil {
		return false, nil
	}

	return b.leftNode.ValueExists(val)
}
