package response

type TransactionResponse struct {
	Id         int              `json:"id"`
	Properties PropertyResponse `json:"Properties"`
	Amount     float64          `json:"amount"`
	Status     string           `json:"status"`
	DateOffer  string           `json:"dateOffer"`
}
