package go_separate

import (
	"strconv"
)

func FormatLargeNumber(numberToFormat float64) string {
	numberAsString := strconv.FormatFloat(numberToFormat, 'f', 0, 64)

	for i := len(numberAsString) - 3; i > 0; i-=3 {
		numberAsString = insertCommaAt(i, numberAsString)
	}
	return numberAsString
}

func insertCommaAt(indexToAddComma int, str string) string {
	return str[:indexToAddComma] + "," + str[indexToAddComma:]
}
