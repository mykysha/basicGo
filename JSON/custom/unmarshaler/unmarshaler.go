package unmarshaler

import (
	"errors"
	"strconv"
)

const (
	stringType  = "string"
	numberType  = "float64"
	booleanType = "bool"
)

const (
	booleanTrueString  = "true"
	booleanFalseString = "false"
)

const (
	doubleQuoteSymbol        = '"'
	leftArrayBracketSymbol   = '['
	rightArrayBracketSymbol  = ']'
	leftObjectBracketSymbol  = '{'
	rightObjectBracketSymbol = '}'
	commaSymbol              = ','
)

var (
	ErrNotPossibleToMarshal     = errors.New("given byte array is not possible to marshal")
	errUnexpectedSquareBracket  = errors.New("unexpected \"]\"")
	errUnexpectedFigureeBracket = errors.New("unexpected \"}\"")
	errBracketNotClosed         = errors.New("brackets were opened, but never closed")
	errTrailingComma            = errors.New("unexpected comma at the end")
)

func Unmarshal(data []byte) (res interface{}, resType string, err error) {
	unicodeData := convertToRune(data)

	switch {
	case unicodeData[0] == doubleQuoteSymbol && unicodeData[len(unicodeData)-1] == doubleQuoteSymbol:
		return stringUnmarshal(unicodeData), stringType, nil
	case isNumber(unicodeData):
		return numberUnmarshal(unicodeData), numberType, nil
	case string(unicodeData) == booleanTrueString || string(unicodeData) == booleanFalseString:
		return booleanUnmarshal(unicodeData), booleanType, nil
	}

	return nil, "", ErrNotPossibleToMarshal
}

func stringUnmarshal(data []rune) string {
	s := string(data[1 : len(data)-1])

	return s
}

func numberUnmarshal(data []rune) float64 {
	bSize := 64

	n, _ := strconv.ParseFloat(string(data), bSize)

	return n
}

func booleanUnmarshal(data []rune) bool {
	return string(data) == booleanTrueString
}
