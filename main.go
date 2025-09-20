package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


const httpAddress = ":8080"

func main(){
	fmt.Println("Server is running on port", httpAddress)

	server := gin.New()

	server.GET("/health", healthHandler)

	log.Fatal(server.Run(httpAddress))
}

func healthHandler(ctx *gin.Context){
	ctx.String(http.StatusOK, "Healthy!")
}