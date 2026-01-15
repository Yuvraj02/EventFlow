package models

type PaymentRequest struct{
	EventData Event	`json:"event"`
	Data PaymentData `json:"data"`
}