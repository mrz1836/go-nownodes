package main

import (
	"context"
	"log"
	"os"

	"github.com/mrz1836/go-nownodes"
)

func main() {
	c := nownodes.NewClient(nownodes.WithAPIKey(os.Getenv("NOW_NODES_API_KEY")))
	info, err := c.GetMempoolEntry(
		context.Background(), nownodes.BCH, "827325a93582ef70e945db95e87d7aca96a325b63aa5d8f8cebc8c45bd71dd01", "unique-id",
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("found tx in mempool, time: ", info.Result.Time)
}
