package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payments []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		//scriptedPan := ScriptPAN(card)
		payment := types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		}
		payments = append(payments, payment)
	}
	return payments
}

//func ScriptPAN(card types.Card) string {
//	pan := strings.Split(string(card.PAN), " ")
//
//	pan[1] = "xxxx"
//	pan[2] = "xxxx"
//	scriptPAN := pan[0] + " " + pan[1] + " " + pan[2] + " " + pan[3]
//	return scriptPAN
//}
