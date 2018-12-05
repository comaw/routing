package main

import (
	"fmt"
	"net/http"
	"routing/collection"
)

func main() {
	http.HandleFunc("/", collection.AddCollection)
	err := http.ListenAndServe(":9260", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Service start")
}
