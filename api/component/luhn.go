package component

import (
	"fmt"
	"strconv"
)

func IsValidLuhn(number string) bool {
	var validLastDigit bool

	if !isNumeric(number) {
		fmt.Println("ERROR: The input contains non-numeric characters!")
		return false
	}

	lastDigit, _ := strconv.Atoi(string(number[len(number)-1]))
	total := 0

	for i := len(number) - 2; i >= 0; i-- {
		sum := 0
		digit := int(number[i] - '0')

		if (len(number)-i)%2 == 0 {
			digit *= 2
		}

		sum = digit/10 + digit%10
		total += sum
	}

	if total%10 != 0 {
		validLastDigit = (10 - total%10) == lastDigit
		fmt.Println(number, "is a valid credit card number!")
		return validLastDigit
	} else {
		validLastDigit = lastDigit == 0
		fmt.Println(number, "is an invalid credit card number!")
		return validLastDigit
	}
}
