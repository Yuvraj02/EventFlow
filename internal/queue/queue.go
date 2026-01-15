package queue

import "eventflow/internal/models"

var EventBuffer chan models.PaymentRequest = make(chan models.PaymentRequest, 2)