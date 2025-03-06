package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	initalBalance := 500.0
	acBank := BankAccount{Balance: initalBalance}

	t.Run("Withdraw money", func(t *testing.T) {
		balanceAmount := acBank.Balance
		withdrawAmount := 100.00
		newbalance := balanceAmount - withdrawAmount
		finalbalance, err := acBank.Withdraw(withdrawAmount)

		assert.NoError(t, err)
		assert.Equal(t, "Balance Amount", newbalance, finalbalance)

	})

	t.Run("Morethan balance", func(t *testing.T) {
		withdrawAmount := 600.00
		initalBalance := acBank.Balance
		finalbalance, err := acBank.Withdraw(withdrawAmount)

		assert.Error(t, err)
		assert.Equal(t, initalBalance, finalbalance)

	})

	t.Run("Zero Withdraw", func(t *testing.T) {
		withdrawBalance := 0.00
		initalBalance := acBank.Balance
		finalbalance, err := acBank.Withdraw(withdrawBalance)

		assert.Error(t, err)
		assert.Equal(t, initalBalance, finalbalance)
	})
}
