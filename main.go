package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/willamesSantoos/simple-api-go/internal/routes"
)

func main() {
	r := mux.NewRouter()
	routes.InitializeRoutes(r)
	
	err := http.ListenAndServe(":3000", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
