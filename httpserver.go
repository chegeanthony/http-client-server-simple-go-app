package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}
