package test

import "errors"

func Add(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("neither a or b is zero")
	}
	return a + b, nil
}
