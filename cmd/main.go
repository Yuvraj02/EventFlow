package main

import (
	"context"
	"eventflow/internal/router"
	"eventflow/internal/workers"
	"fmt"
	"net/http"
)


func main(){

	mux:=http.NewServeMux()
	api.Router(mux)

	fmt.Println("Server Started on Port 8080")
	var workCount int = 2
	for i := 0; i < workCount; i++ {
		go workers.StartPaymentWorker(context.Background())
	}

	var err error = http.ListenAndServe(":8080", mux)
	
	if err!=nil{
		fmt.Println(err)
	}

}