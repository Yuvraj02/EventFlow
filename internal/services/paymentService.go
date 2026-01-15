package services

import (
	"context"
	"eventflow/internal/models"
	"eventflow/internal/storage"
	"time"
)

func ProcessPayment(ctx context.Context, pay *models.PaymentData){
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.PaymentData[storage.PaymentKey] = *pay
	storage.PaymentKey++
	time.Sleep(3*time.Second)
}
