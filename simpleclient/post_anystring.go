package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main(){

	reader := strings.NewReader("テキスト")

	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}