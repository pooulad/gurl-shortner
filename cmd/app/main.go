package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pooulad/gurl-shortner/pkg"
)

func main() {
	shortener := pkg.NewUrlShortner(make(map[string]string))

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	host := os.Getenv("HOST_ADDR")
	port := os.Getenv("HOST_PORT")

	http.HandleFunc("/", shortener.HandleRoot)
	http.HandleFunc("/shorten", shortener.HandleShortenUrl)
	http.HandleFunc("/short/", shortener.HandleRedirectUrl)

	fmt.Printf("gurl-shortner is running on %v:%v\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil)
	if err != nil {
		panic(err)
	}
}
