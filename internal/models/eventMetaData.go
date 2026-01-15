package models

type Event struct {
	EventId string 	`json:"event_id"`
	Name string		`json:"name"`
	Timestamp int	`json:"timestamp"`
	PaymentId string `json:"payment_id"`
}