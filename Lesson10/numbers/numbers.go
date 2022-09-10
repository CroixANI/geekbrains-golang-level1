package numbers

import "errors"

var (
	ErrorIncorrectNumber = errors.New("this function works with only three digits number")
)

func SplitNumber(number int) (int, int, int, error) {
	if number > 999 || number <= 99 {
		return 0, 0, 0, ErrorIncorrectNumber
	}

	var hundreds = number / 100
	var tens = (number - (hundreds * 100)) / 10
	var units = number - (hundreds * 100) - (tens * 10)
	return units, tens, hundreds, nil
}
