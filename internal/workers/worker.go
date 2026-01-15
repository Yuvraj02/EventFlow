package workers

import (
	"context"
	"eventflow/internal/queue"
	"eventflow/internal/services"
	"time"
)

func StartPaymentWorker(ctx context.Context){

	for {
		select {
		case pay := <-queue.EventBuffer:
			paymentCtx,cancel := context.WithTimeout(ctx, time.Second * 5)
			services.ProcessPayment(paymentCtx, &pay.Data)
			cancel()
		case <- ctx.Done():
			return
		}
	}
}