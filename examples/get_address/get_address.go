package main

import (
	"context"
	"log"
	"os"

	"github.com/mrz1836/go-nownodes"
)

func main() {
	c := nownodes.NewClient(nownodes.WithAPIKey(os.Getenv("NOW_NODES_API_KEY")))
	info, err := c.GetAddress(
		context.Background(), nownodes.BSV, "1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL",
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("found address: ", info.Address, "in txs", info.Txs)
}
