package main

import (
	"fmt"
	"log"
	"net/http"
)


const httpAddress = ":8080"

func main(){
	fmt.Println("Server is running on port", httpAddress)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(httpAddress, mux))
}

func healthHandler(w http.ResponseWriter, _ *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy!"))
}