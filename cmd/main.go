package main

import (
	"eventflow/internal/router"
	"fmt"
	"net/http"
)


func main(){

	mux:=http.NewServeMux()
	api.Router(mux)

	fmt.Println("Server Started on Port 8080")

	var err error = http.ListenAndServe(":8080", mux)
	
	if err!=nil{
		fmt.Println(err)
	}

}