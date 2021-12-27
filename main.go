package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Set routing rules
	http.HandleFunc("/", Tmp)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("DB_PORT")
	log.Println(port)
	io.WriteString(w, "version 1")
}
