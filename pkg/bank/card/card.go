package card

import (
	"bank/pkg/bank/types"
)

func Total(cards []types.Card) types.Money {
	sum := int64(0)
	for _, operation := range cards {
		if !operation.Active {
			continue
		}
		if operation.Balance <= 0 {
			continue
		}
		sum += int64(operation.Balance)
	}
	return types.Money(sum)
}
