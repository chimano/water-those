package main

import (
	"fmt"

	"github.com/chimano/water-those-service/api"
)

func main() {
	srv, err := api.NewServer()

	if err != nil {
		fmt.Println("ouch")
	}

	srv.Start()
}
