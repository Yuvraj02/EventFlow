package storage

import (
	"eventflow/internal/models"
	"sync"
)

//Temporary storage
var PaymentData map[int]models.PaymentData
var PaymentKey int
var Mutex sync.Mutex

func init(){
	PaymentData = make(map[int]models.PaymentData)
	PaymentKey = 0
}