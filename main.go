package main

import (
	"context"
	"flag"
	"log"

	"github.com/railgun-0402/DI-Golang/app/cmd"
)

func main() {
	mode := flag.String("mode", "server", "Mode: server or worker")
	flag.Parse()

	ctx := context.Background()

	switch *mode {
	case "server":
		e := cmd.NewAPI()
		log.Fatal(e.Start(":8080"))
	case "worker":
		cmd.RunWorker(ctx)
	default:
		log.Fatalf("unknown mode: %s", *mode)
	}
}