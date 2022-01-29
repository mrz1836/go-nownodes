package main

import (
	"context"
	"log"
	"os"

	"github.com/mrz1836/go-nownodes"
)

func main() {
	c := nownodes.NewClient(nownodes.WithAPIKey(os.Getenv("NOW_NODES_API_KEY")))
	info, err := c.GetTransaction(
		context.Background(), nownodes.BSV, "17961a51337369bf64e45e8410a7ce4cfb0c88b5d883d9e8a939dfdd0f7591fd",
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("found tx: ", info.TxID, "in block", info.BlockHash)
}
