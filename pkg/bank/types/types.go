package types

//Моней представляет собой денежную сумму в минимальных еденицы(центы, копеек, дирамов и.т.д)

//Payment предстовляет собой инфо о платеже

type Payment struct {
	ID     int
	Amount Money
}
type Money int64

//Currency предстоаляет код валюты
type Currency string

//Коды валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет infotmation by Card
type Card struct {
	ID         int
	PAN        PAN
	MinBalance Money
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}
