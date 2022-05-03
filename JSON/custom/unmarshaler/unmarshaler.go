package unmarshaler

import (
	"errors"
	"strconv"
)

const (
	stringType  = "string"
	numberType  = "float64"
	booleanType = "bool"
	arrayType   = "array of interfaces"
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

var ErrNotPossibleToMarshal = errors.New("given byte array is not possible to marshal")

func Unmarshal(data []byte) (res interface{}, resType string, err error) {
	unicodeData := convertToRune(data)

	firstEl := unicodeData[0]
	lastEl := unicodeData[len(unicodeData)-1]

	switch {
	case firstEl == leftArrayBracketSymbol && lastEl == rightArrayBracketSymbol:
		res, err := arrayUnmarshal(unicodeData)

		return res, arrayType, err
	case firstEl == doubleQuoteSymbol && lastEl == doubleQuoteSymbol:
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
