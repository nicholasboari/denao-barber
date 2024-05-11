package main

import (
	"log"

	"github.com/nicholasboari/denao-barber/configs"
)

func main() {
	err := configs.Init()
	if err != nil {
		log.Default().Fatalf("config initialization error: %v", err)
		return
	}
}
