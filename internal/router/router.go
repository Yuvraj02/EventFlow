package api

import (
	"eventflow/internal/handlers"
	"net/http"
)

func Router(mux *http.ServeMux){

	mux.HandleFunc("/", handlers.BasicGetReq)
	mux.HandleFunc("/hello", handlers.HelloWorld)
	mux.HandleFunc("/register", handlers.RegisterUser)
	mux.HandleFunc("/pay", handlers.PaymentRequest)
	mux.HandleFunc("/get_pd", handlers.GetPaymentData)
}	