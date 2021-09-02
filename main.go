package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/mabna/service"
)

func main() {
	Cli()

	service.Start()
}
