package main

import (
	"fmt"
	"log"

	"github.com/yeager-eren/rango-sdk-go/pkg/client"
)

type Config struct {
	device_id string
	api_key   string
	api_url   string
}

func main() {

	cfg := client.Config{
		DeviceID: "gosdk-uuid-uuid-uuid",
		APIKey:   "4a624ab5-16ff-4f96-90b7-ab00ddfc342c",
		APIURL:   "https://api.rango.exchange",
	}

	client := client.New(cfg)
	// chains, err := client.Chains()
	// meta, err := client.Meta()
	// swappers, err := client.Swappers()
	messagingProtocols, err := client.MessagingProtocols()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: %d", messagingProtocols)
	// fmt.Println("Result: %d", len(chains))
	// fmt.Println("First one: %s", chains[0].Name)
	// fmt.Println("---------------------------")
	// fmt.Println(chains)
}
