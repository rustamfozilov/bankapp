package payment

import (
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {
	maxPayment := payments[0]
	for _, payment := range payments {
		if payment.Amount >
			maxPayment.Amount {
			maxPayment = payment
		}
	}
	return maxPayment
}
