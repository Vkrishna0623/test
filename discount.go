package test

import "errors"

func CalculateDiscount(price float64, discount float64) (float64, error) {
	if price < 0 || discount < 0 {
		return 0, errors.New("price and discount cannot be negative")
	}
	if discount > 100 {
		return 0, errors.New("discount cannot be greater than 100%")
	}
	return price - (price * discount / 100), nil
}
