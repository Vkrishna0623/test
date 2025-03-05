package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessOrder(t *testing.T) {
	t.Run("Invalid Amount", func(t *testing.T) {
		order := Order{Amount: 0.0, Paid: false}

		res, err := ProcessOrder(order)

		assert.Error(t, err)
		assert.Equal(t, "Amount should be grater than 0", err.Error())
		assert.Equal(t, "", res)
	})

	t.Run("Paid Amount", func(t *testing.T) {
		order := Order{Amount: 100, Paid: true}
		res, err := ProcessOrder(order)

		assert.Error(t, err)
		assert.Equal(t, "AMount Paid Successfully", res)
	})

	t.Run("Order Placed", func(t *testing.T) {
		order := Order{Amount: 100, Paid: false}

		res, err := ProcessOrder(order)

		assert.Error(t, err)
		assert.Equal(t, "Order Placed Successfully", res)
	})
}
