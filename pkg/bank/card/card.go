package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payments []types.PaymentSource
	for _, card := range cards {
		if !card.Active || card.Balance <= 0 {
			continue
		}
		payments = append(payments, types.PaymentSource{
			Card:    "card",
			Number:  "5555 6666 7777 8888",
			Balance: card.Balance,
		})
	}

	return payments
}
