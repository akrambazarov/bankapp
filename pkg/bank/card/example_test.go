package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	fmt.Println(PaymentSources([]types.Card{
		{
			Balance: 1_000_00,
			PAN:     "5555 6666 7777 8888",
			Active:  true,
		},
	}))
	fmt.Println(PaymentSources([]types.Card{
		{
			Balance: 1_000_00,
			PAN:     "5678 7889 0001 7898",
			Active:  false,
		},
	}))
	fmt.Println(PaymentSources([]types.Card{
		{
			Balance: -1_000_00,
			PAN:     "9085 9977 6678 9087",
			Active:  true,
		},
	}))
	//Output:
	//[{card 5555 6666 7777 8888 100000}]
	//[]
	//[]
}
