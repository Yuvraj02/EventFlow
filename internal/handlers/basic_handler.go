package handlers

import (
	"encoding/json"
	"eventflow/internal/models"
	"eventflow/internal/queue"
	"eventflow/internal/storage"
	"fmt"
	"net/http"
)

func BasicGetReq(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	_ = ctx 

	response:= struct{
		Message string `json:"message"` //As NewEncoder only encodes to JSON with exported fields
	}{
		Message: "This is a Basic Get Request",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err!=nil{
		fmt.Println("Error in basic Get Request: ", err)
		return
	}
}


func HelloWorld(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	_ = ctx

	response := struct {
		Message string `json:"message"`
	}{
		Message:"Hello World",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Not Working Properly : ", err)
		return

	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	_ = ctx

	//Recieve Event	
	//TODO : Implement This later

}

func PaymentRequest(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	_ = ctx

	//Event to be recieve
	var paymentRequest models.PaymentRequest
	json.NewDecoder(r.Body).Decode(&paymentRequest)

	select {
		case queue.EventBuffer <- paymentRequest:
			w.WriteHeader(http.StatusAccepted)
		case <- ctx.Done():
			return 
		default:
			w.WriteHeader(http.StatusTooManyRequests)
	}
}

func GetPaymentData(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	_ = ctx

	var paymentData []models.PaymentData
	
	for _,val := range storage.PaymentData {
		paymentData = append(paymentData, val)
	}

	response := struct{
		Status string `json:"status"`
		Data []models.PaymentData `json:"data"`
	}{
		Status: "success",
		Data: paymentData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}