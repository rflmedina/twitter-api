package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running at port: 5000 | Have fun! :D")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
