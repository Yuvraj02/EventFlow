package models

type PaymentData struct {

	PaymentId string `json:"payment_id"`
	AccountId string `json:"account_id"`
	Ammount string `json:"amount"`
	Currency string `json:"currency"`
	IncludeTax bool `json:"include_tax"`
	Method string `json:"method"`
	Status string `json:"status"`
}