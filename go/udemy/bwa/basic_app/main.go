package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8081"

func main() {
	fmt.Println("start")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println("HTTP creation error", err)
		}

		fmt.Printf("%d bytes written", n)
	})

	_ = http.ListenAndServe(":8081", nil)

}
