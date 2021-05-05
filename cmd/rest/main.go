package main

import (
	"fmt"
	"github.com/sfchong/gorest/http"
	"github.com/sfchong/gorest/mock"
	"log"
)

func main() {
	fmt.Println("Initializing http server...")
	s := http.NewServer()

	fmt.Println("Attaching service to server...")
	s.UserService = mock.NewUserService()

	fmt.Println("Listen and Serve...")
	err := s.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
