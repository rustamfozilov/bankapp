package card

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

func ExampleMax2() {
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
			ID:     67,
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
	// 67
}

//func ExampleScriptPAN() {
//	card := types.Card{
//		ID:         0,
//		PAN:        "5555 6756 6776 9999",
//		Balance:    50,
//		Currency:   "TJS",
//		Color:      "white",
//		Name:       "card",
//		Active:     true,
//		MinBalance: 0,
//	}
//
//	fmt.Println(ScriptPAN(card))
//	// Output:
//	// 5555 xxxx xxxx 9999
//}

func ExamplePaymentSources() {
	card := types.Card{
		ID:         0,
		PAN:        "5555 6756 6776 9999",
		Balance:    50,
		Currency:   "TJS",
		Color:      "white",
		Name:       "card",
		Active:     true,
		MinBalance: 0,
	}
	card2 := types.Card{
		ID:         0,
		PAN:        "5555 6756 6776 6666",
		Balance:    666,
		Currency:   "TJS",
		Color:      "white",
		Name:       "card",
		Active:     true,
		MinBalance: 0,
	}

	cards := []types.Card{card, card2}

	payments := PaymentSources(cards)
	for _, p := range payments {
		fmt.Println(p.Number)
	}
// Output:
// 5555 6756 6776 9999
// 5555 6756 6776 6666
}
