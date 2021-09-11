package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     99,
			Amount: 23,
		},
		{
			ID:     24,
			Amount: 45,
		},
		{
			ID:     66,
			Amount: 555,
		},
		{
			ID:     3,
			Amount: 555,
		},
		{
			ID:     9,
			Amount: 267,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	// Output:
	// 66
}
