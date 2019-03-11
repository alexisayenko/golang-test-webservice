package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Yo yo yo yo!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Println(text)
}
