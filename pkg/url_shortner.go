package pkg

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type UrlShortner struct {
	urls map[string]string
}

func NewUrlShortner(urls map[string]string) *UrlShortner {
	return &UrlShortner{
		urls: urls,
	}
}

func (u *UrlShortner) HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	  <head>
		<title>gurl-shortner</title>
		<link rel="stylesheet" href="/static/index.css" />
	  </head>
	  <body>
		<main class="main">
		  <h2>gurl-shortner</h2>
		  <form method="post" action="/shorten">
			<input
			  class="input"
			  type="url"
			  name="url"
			  placeholder="Enter a URL"
			  required
			/>
			<input class="btn" type="submit" value="Make short now" />
		  </form>
		</main>
	  </body>
	</html>
	
		`)
}

func (u *UrlShortner) HandleShortenUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	shortKey := shortnerKeyGenerator()
	u.urls[shortKey] = originalURL

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	host := os.Getenv("HOST_ADDR")
	port := os.Getenv("HOST_PORT")

	shortenedURL := fmt.Sprintf("http://%v:%v/short/%s", host, port, shortKey)

	w.Header().Set("Content-Type", "text/html")
	responseHTML := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>gurl-shortner</title>
		<link rel="stylesheet" href="/static/index.css" />
	</head>
	<body>
       <main class="main">
	   <h2>gurl-shortner</h2>
	   <p>Original URL: %s</p>
	   <p>Shortened URL: <a target="_blank" href="%s">%s</a></p>
	   <form method="post" action="/shorten">
			<input
			  class="input"
			  type="url"
			  name="url"
			  placeholder="Enter a URL"
			  required
			/>
			<input class="btn" type="submit" value="Make short now" />
		  </form>
	   </main>
		</body>
		</html>
    `, originalURL, shortenedURL, shortenedURL)
	fmt.Fprint(w, responseHTML)
}
func (u *UrlShortner) HandleRedirectUrl(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	originalURL, found := u.urls[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func shortnerKeyGenerator() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
