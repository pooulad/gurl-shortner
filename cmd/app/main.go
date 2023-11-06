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
