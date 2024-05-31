package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/sharma03r/orders-api/application"
)

func main() {
	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel() //should be called at the end of the current function

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start the server:", err)
	}
}
