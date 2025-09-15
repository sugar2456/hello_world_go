package calc

import (
	"errors"
)

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return 0, 0, errors.New("0で割ることはできません")
	}
	result = num / denom
	remainder = num % denom
	return result, remainder, nil
}
