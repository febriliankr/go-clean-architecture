package main

import (
	"log"

	"github.com/febriliankr/go-clean-architecture/internal/app"
	"github.com/febriliankr/go-clean-architecture/internal/delivery/rest"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file, will read directly from ENV vars instead")
	}

	// init app
	ShipmentApp, err := app.NewShipmentApp()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		errs := ShipmentApp.Close()
		for _, e := range errs {
			if e != nil {
				log.Println(e)
			}
		}
	}()

	// Start REST
	rest.Start(ShipmentApp)
}
