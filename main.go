package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("wwwroot")))

	log.Println("Listening at port 5001")
	log.Fatalln(http.ListenAndServe(":5001", nil))
}
