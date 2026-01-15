package services

import (
	"eventflow/internal/models"
	"eventflow/internal/storage"
)

func ProcessPayment(pay *models.PaymentData){
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.PaymentData[storage.PaymentKey] = *pay
	storage.PaymentKey++
}
