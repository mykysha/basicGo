package binarytree

import (
	"fmt"
)

type BinaryTree interface {
	Insert(interface{}) error
}

type BinaryNode struct {
	treeType  string
	value     interface{}
	leftNode  *BinaryNode
	rightNode *BinaryNode
}

func NewBinaryTree() *BinaryNode {
	return &BinaryNode{
		treeType:  "",
		value:     nil,
		leftNode:  nil,
		rightNode: nil,
	}
}

func newBinaryNode(val interface{}, t string) *BinaryNode {
	return &BinaryNode{
		treeType:  t,
		value:     val,
		leftNode:  nil,
		rightNode: nil,
	}
}

func (b *BinaryNode) Insert(newValue interface{}) error {
	switch v := newValue.(type) {
	case bool:
		err := b.insertBool(v)
		if err != nil {
			return fmt.Errorf("bool insert: %w", err)
		}

		return nil
	case string:
		err := b.insertString(v)
		if err != nil {
			return fmt.Errorf("string insert: %w", err)
		}

		return nil
	case int:
		err := b.insertInt(v)
		if err != nil {
			return fmt.Errorf("int insert: %w", err)
		}

		return nil
	case rune:
		err := b.insertRune(v)
		if err != nil {
			return fmt.Errorf("rune insert: %w", err)
		}

		return nil
	case float64:
		err := b.insertFloat64(v)
		if err != nil {
			return fmt.Errorf("float64 insert: %w", err)
		}

		return nil
	default:
		return errUnknownType
	}
}

func (b *BinaryNode) ValueExists(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return b.boolExists(v)
	case string:
		return b.stringExists(v)
	case int:
		return b.intExists(v)
	case rune:
		return b.runeExists(v)
	case float64:
		return b.float64Exists(v)
	default:
		return false, errUnknownType
	}
}

func (b *BinaryNode) GetMaxValue() interface{} {
	if b.rightNode == nil {
		return b.value
	}

	return b.rightNode.GetMaxValue()
}

func (b *BinaryNode) GetMinValue() interface{} {
	if b.leftNode == nil {
		return b.value
	}

	return b.leftNode.GetMinValue()
}

func (b *BinaryNode) GetInOrder() []interface{} {
	var order []interface{}

	if b.leftNode != nil {
		order = append(order, b.leftNode.GetInOrder()...)
	}

	order = append(order, b.value)

	if b.rightNode != nil {
		order = append(order, b.rightNode.GetInOrder()...)
	}

	return order
}
