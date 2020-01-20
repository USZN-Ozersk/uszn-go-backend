package main

import "github.com/USZN-Ozersk/uszn-go-backend/internail/app/apiserver"

import "log"

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
