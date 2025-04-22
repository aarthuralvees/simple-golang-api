package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aarthuralvees/simple-go-api/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting api...")
	handlers.Handler(router)
	fmt.Println("runing on localhost 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
