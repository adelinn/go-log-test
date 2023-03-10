package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	// Hello world, the web server
	log.Println("Starting application...")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
