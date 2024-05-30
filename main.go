package main

import (
	"fmt"
	"context"
	"github.com/sharma03r/orders-api/application"
)
func main(){
	app := application.New()
	err := app.Start(context.TODO())

	if err != nil{
		fmt.Errorf("failed to start the server: %w", err)
	}
}