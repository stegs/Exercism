package romannumerals

import (
	"errors"
)

//ToRomanNumeral accepts an integer number and converts it into a roman numeral
func ToRomanNumeral(number int) (string, error) {

	if number <= 0 || number > 3000 {
		return "", errors.New("Invalid number")
	}

	roman := ""
	stringToAdd := ""
	times := 0

	for number > 0 {
		times = 0
		if number >= 1000 {
			times = number / 1000
			number -= times * 1000
			stringToAdd = "M"

		} else if number >= 500 {
			if number > 900 {
				roman += "CM"
				number -= 900
			} else {
				roman += "D"
				number -= 500
			}

		} else if number >= 100 {
			if number > 400 {
				roman += "CD"
				number -= 400
			} else {
				times = number / 100
				number -= times * 100
				stringToAdd = "C"
			}

		} else if number >= 50 {
			if number > 90 {
				roman += "XC"
				number -= 90
			} else {
				roman += "L"
				number -= 50
			}

		} else if number >= 10 {
			if number > 40 {
				roman += "XL"
				number -= 40
			} else {
				times = number / 10
				number -= times * 10
				stringToAdd = "X"
			}

		} else if number >= 5 {
			if number == 9 {
				roman += "IX"
				number -= 9
			} else {
				roman += "V"
				number -= 5
			}

		} else if number >= 1 {
			if number == 4 {
				roman += "IV"
				number -= 4
			} else {
				times = number
				number = 0
				stringToAdd = "I"
			}
		}

		for times > 0 {
			roman += stringToAdd
			times--
		}
	}

	return roman, nil
}
