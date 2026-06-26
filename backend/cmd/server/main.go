package main

import (
	server "github.com/9inejames/pok-deng/backend/internal/httpapi"
)

func main() {
	server.NewFiberServer("8080").Start()
}
