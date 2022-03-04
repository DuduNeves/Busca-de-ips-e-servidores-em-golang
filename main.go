package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	aplicação := app.Gerar()
	err := aplicação.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
