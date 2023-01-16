package main

import (
	"fmt"

	"sentiment/api"
)

func main() {

	endpointData := api.Endpoint("http://127.0.0.1:8000/")
	fmt.Println("Endpoint: ", endpointData)

	// api.Sentiment()
}
