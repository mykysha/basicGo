package unmarshaler

import (
	"unicode"
	"unicode/utf8"
)

func isNumber(r []rune) bool {
	dotCounter := 0

	for _, el := range r {
		if !unicode.IsDigit(el) {
			if el == '.' && dotCounter == 0 {
				dotCounter += 1

				continue
			}

			return false
		}
	}

	return true
}

func convertToRune(data []byte) []rune {
	var (
		res   []rune
		total int
	)

	dataString := string(data)

	dataLen := len(data)

	for {
		symbol, runeCount := utf8.DecodeRuneInString(dataString)

		symbol = unicode.ToLower(symbol)

		res = append(res, symbol)

		total += runeCount

		if total == dataLen {
			break
		}

		dataString = dataString[runeCount:]
	}

	return res
}
