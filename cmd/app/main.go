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

	http.HandleFunc("/shorten", shortener.HandleShorten)
	http.HandleFunc("/short/", shortener.HandleRedirect)

