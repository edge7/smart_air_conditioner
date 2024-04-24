package main

import (
	"log"

	py "air_driver/interfaces_to_python"
)

func main() {
	err, temperature := py.GetTemperature()

	if err != nil {
		log.Fatalf("could not get model result: %v", err)
	}

	log.Println("Temperature:", temperature)

	err, modelResult := py.GetModelResult("input_path")

	if err != nil {
		log.Fatalf("could not get model result: %v", err)
	}
	log.Println("Model result:", modelResult)
}
