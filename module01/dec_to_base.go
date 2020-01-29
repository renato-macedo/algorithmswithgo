package module01

import (
	"fmt"
	"strconv"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {

	res := ""
	rem := 0
	for dec > 0 {
		rem = dec % base
		if rem > 9 {
			res = fmt.Sprintf("%s%s", convertToLetter(rem), res)
		} else {
			res = fmt.Sprintf("%d%s", rem, res)
		}

		dec = dec / base
	}

	return res
}

func convertToLetter(number int) string {
	switch number {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return strconv.Itoa(number)
	}
}
