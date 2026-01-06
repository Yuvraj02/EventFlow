package handlers

import (
	"encoding/json"
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

