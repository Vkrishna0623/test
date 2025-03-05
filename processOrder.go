package test

import "errors"

type Order struct {
	Amount float64
	Paid   bool
}

func ProcessOrder(order Order) (string, error) {
	if order.Amount <= 0 {
		return "", errors.New("order amount must be greater than zero")
	}
	if order.Paid {
		return "Order already paid", nil
	}
	return "Order processed successfully", nil
}
