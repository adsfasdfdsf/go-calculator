package main

import (
	"go-calculator/internal/endpoint/server"
	"go-calculator/internal/repository/calc"
)

const (
	SERVER_PORT = 8080
)

func main() {
	repo := calc.Calculator{}
	srv := server.NewServer(SERVER_PORT, repo)
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
