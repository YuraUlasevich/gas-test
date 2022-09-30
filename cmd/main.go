package main

import (
	"errors"
	"fmt"
	gas2 "gas-test/internal/adapters/gas"
	"gas-test/internal/domain/gas"
	"net/http"
	"os"
)

func main() {
	gasService := gas.NewService()
	gasHandlers := gas2.NewHandler(gasService)

	http.HandleFunc("/", gasHandlers.Statistic)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
