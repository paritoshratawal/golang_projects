package main

import (
	"fmt"
)


func main(){
	fmt.Println("Welcome to the Bank")
	server := NewAPIServer(":8080")
	server.Run()
}