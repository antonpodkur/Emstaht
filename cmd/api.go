package main

import (
	"github.com/antonpodkur/Emstaht/internal/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
