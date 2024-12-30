package main

import (
	"carbon-tax-ledger/app/controllers"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	for {
		mintedAmount, err := controllers.MintCarbonToken()
		if err != nil {
			log.Println("Failed to mint carbon token:", err)
		} else {
			log.Println("Minted", mintedAmount, "carbon token(s)")
		}

		time.Sleep(3 * time.Second)
	}
}
