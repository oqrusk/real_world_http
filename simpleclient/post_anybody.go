package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func main(){

	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}