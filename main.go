// basic go web server that will listen on port 8080

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.ListenAndServe( ":8080", nil)
}