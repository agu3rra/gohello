package main

import (
	"fmt"
	"net/http"
)

const port = "8081"

func helloHandler(response http.ResponseWriter, request *http.Request) {
	user := request.URL.Query().Get("user")
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if user == "" {
		fmt.Fprint(response, "Hello stranger!")
		return
	}
	fmt.Fprintf(response, "Hello %s!", user)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Serving on port %s. Hit CTRL+C to stop.\n", port)
	http.ListenAndServe(":"+port, nil)
}
