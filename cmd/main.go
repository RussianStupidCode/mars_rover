package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"rover/pkg/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	err := application.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
