package payment

import (
	"bank/pkg/bank/types"
	"strings"
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


func PaymentSources(cards []types.Card) []types.PaymentSource {
	 var payments []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0{
			continue
		}
		scriptedPan := ScriptPAN(card)
		payment := types.PaymentSource{
			Type:    "card",
			Number: scriptedPan ,
			Balance: card.Balance,
		}
		payments = append(payments, payment)
	}
return payments}

func ScriptPAN(card types.Card) string {
	pan := strings.Split(string(card.PAN), " ")
	pan[1] = "xxxx"
	pan[2] = "xxxx"
	scriptPAN := pan[0] + " " + pan[1] + " " + pan[2] + " " + pan[3]
	return scriptPAN
}
