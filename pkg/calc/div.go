package calc

import (
	"errors"
)

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}
