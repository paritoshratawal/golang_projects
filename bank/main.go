package main

import (
	"fmt"
	"log"
)

func main(){
	fmt.Println("Welcome to the Bank")
	store, err := newPostgresStorage()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := store.Init(); err != nil{
		log.Fatalf("Failed to initialize database: %v", err)
	}
	server := NewAPIServer(":8080", store)
	server.Run()
}