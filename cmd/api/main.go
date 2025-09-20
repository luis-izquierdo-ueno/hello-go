package main

import (
	"hello-go/cmd/api/bootstrap"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){
	
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
