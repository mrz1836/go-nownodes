package main

import (
	"context"
	"log"
	"os"

	"github.com/mrz1836/go-nownodes"
)

func main() {
	c := nownodes.NewClient(nownodes.WithAPIKey(os.Getenv("NOW_NODES_API_KEY")))
	result, err := c.SendTransaction(
		context.Background(), nownodes.BSV, "010000000192c4fb43a78e0b44b7825f8...",
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("send success: ", result.Result)
}
