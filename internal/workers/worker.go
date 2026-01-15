package workers

import (
	"eventflow/internal/queue"
	"eventflow/internal/services"
	"fmt"
)

func PaymentWorker(){
	select {
	case pay :=  <-queue.EventBuffer:
			services.ProcessPayment(&pay.Data)
	default:
		fmt.Println("No Event to Process")
	}

}