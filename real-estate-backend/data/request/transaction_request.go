package request

type TransactionCreateRequest struct {
	PropertyId int
	BuyerId    int
	SellerId   int
	Amount     float64
	Status     string
	DateOffer  string
}
