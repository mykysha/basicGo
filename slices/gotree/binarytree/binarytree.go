package binarytree

import (
	"errors"
	"strings"
)

var ErrNodeExists = errors.New("node with this value already exists")

func NewBinaryNode(val string) *BinaryNode {
	return &BinaryNode{
		value:     val,
		leftNode:  nil,
		rightNode: nil,
	}
}

type BinaryNode struct {
	value     string
	leftNode  *BinaryNode
	rightNode *BinaryNode
}

func (b *BinaryNode) Insert(newValue string) error {
	newValue = strings.ToLower(newValue)

	comp := strings.Compare(b.value, newValue)

	if comp == 0 {
		return ErrNodeExists
	}

	if comp == -1 {
		if b.rightNode == nil {
			b.rightNode = NewBinaryNode(newValue)

			return nil
		}

		return b.rightNode.Insert(newValue)
	}

	if comp == 1 {
		if b.leftNode == nil {
			b.leftNode = NewBinaryNode(newValue)

			return nil
		}

		return b.leftNode.Insert(newValue)
	}

	return nil
}

func (b *BinaryNode) ValueExists(value string) bool {
	value = strings.ToLower(value)

	comp := strings.Compare(b.value, value)

	if comp == 0 {
		return true
	}

	if comp == -1 {
		if b.rightNode == nil {
			return false
		}

		return b.rightNode.ValueExists(value)
	}

	if comp == 1 {
		if b.leftNode == nil {
			return false
		}

		return b.leftNode.ValueExists(value)
	}

	return false
}

func (b *BinaryNode) GetMaxValue() string {
	if b.rightNode == nil {
		return b.value
	}

	return b.rightNode.GetMaxValue()
}

func (b *BinaryNode) GetMinValue() string {
	if b.leftNode == nil {
		return b.value
	}

	return b.leftNode.GetMaxValue()
}

func (b *BinaryNode) GetInOrder() []string {
	var order []string

	if b.leftNode != nil {
		order = append(order, b.leftNode.GetInOrder()...)
	}

	order = append(order, b.value)

	if b.rightNode != nil {
		order = append(order, b.rightNode.GetInOrder()...)
	}

	return order
}
