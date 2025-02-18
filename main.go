package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/scores", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Football scores will be here!")
	})

	fmt.Println("Server is listening on port 5050...")
	http.ListenAndServe(":5050", nil)
}
