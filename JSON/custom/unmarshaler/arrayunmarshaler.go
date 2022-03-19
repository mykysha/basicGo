package unmarshaler

import (
	"errors"
	"fmt"
	"unicode"
)

const backslashSymbol = '\\'

var (
	errTrailingComma         = errors.New("unexpected comma at the end")
	errBracketNotClosed      = errors.New("brackets were opened, but never closed")
	errNoCommaAfterElement   = errors.New("found no comma or EOF after array element")
	errDoubleQuotesNotClosed = errors.New("double quotes were opened, but never closed")
	errInvalidEscapeSequence = errors.New("\\ symbol should be accompanied by a valid escape key")
)

type arrayInfo struct {
	read          int
	array         []interface{}
	data          []rune
	currentObject []rune
	expected      []rune
}

func arrayUnmarshal(data []rune) ([]interface{}, error) {
	a := &arrayInfo{
		read:          0,
		array:         nil,
		data:          data,
		currentObject: nil,
		expected:      nil,
	}

	a.trimParentheses()

	if a.checkTrailingComma() {
		return nil, errTrailingComma
	}

	err := a.unmarshal()
	if err != nil {
		return nil, fmt.Errorf("object unmarshal: %w", err)
	}

	return a.array, nil
}

func (a *arrayInfo) trimParentheses() {
	a.data = a.data[1 : len(a.data)-1] // removed start and end brackets
}

// checkTrailingComma returns true if json object ends with trailing comma.
func (a arrayInfo) checkTrailingComma() bool {
	return a.data[len(a.data)-1] == commaSymbol
}

func (a *arrayInfo) unmarshal() error {
	for a.read < len(a.data) {
		symbol := a.data[a.read]

		if unicode.IsSpace(symbol) {
			a.read++

			continue
		}

		switch symbol {
		case doubleQuoteSymbol:
			err := a.handleString()
			if err != nil {
				return fmt.Errorf("string: %w", err)
			}
		case leftArrayBracketSymbol, leftObjectBracketSymbol:
			err := a.handleBracketedExpression()
			if err != nil {
				return fmt.Errorf("bracketed expression: %w", err)
			}
		default:
			err := a.handleSimpleExpression()
			if err != nil {
				return fmt.Errorf("simple type: %w", err)
			}
		}
	}

	return nil
}

func (a *arrayInfo) handleString() error {
	a.currentObject = append(a.currentObject, doubleQuoteSymbol)

	a.read++

	for a.read < len(a.data) {
		symbol := a.data[a.read]

		a.currentObject = append(a.currentObject, symbol)

		a.read++

		switch symbol {
		case backslashSymbol:
			if a.read == len(a.data) {
				return errInvalidEscapeSequence
			}

			symbol = a.data[a.read]
			a.currentObject = append(a.currentObject, symbol)

			switch {
			case !isValidEscapeCharacter(symbol):
				return errInvalidEscapeSequence
			}

		case doubleQuoteSymbol:
			if a.read == len(a.data) || a.data[a.read] == commaSymbol {
				a.read++

				return a.unmarshalElement()
			}

			return errNoCommaAfterElement
		}
	}

	return errDoubleQuotesNotClosed
}

func isValidEscapeCharacter(symbol rune) bool {
	escapeSequenceCharacters := [...]rune{'t', 'n', 'f', '\\', 'b', 'r', '"'}

	for _, el := range escapeSequenceCharacters {
		if symbol == el {
			return true
		}
	}

	return false
}

func (a *arrayInfo) handleBracketedExpression() error {
	return nil
}

func (a *arrayInfo) handleSimpleExpression() error {
	for a.read < len(a.data) {
		symbol := a.data[a.read]

		a.read++

		switch symbol {
		case commaSymbol:
			err := a.unmarshalElement()
			if err != nil {
				return fmt.Errorf("simple type unmarshal: %w", err)
			}
		default:
			a.currentObject = append(a.currentObject, symbol)
		}
	}

	return a.unmarshalElement()
}

func (a *arrayInfo) unmarshalElement() error {
	obj, _, err := Unmarshal([]byte(string(a.currentObject)))
	if err != nil {
		return fmt.Errorf("array element unmarshal: %w", err)
	}

	a.currentObject = nil

	a.array = append(a.array, obj)

	return nil
}
