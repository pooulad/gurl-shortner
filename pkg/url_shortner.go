package pkg

import "net/http"

type UrlShortner struct {
	urls map[string]string
}

func NewUrlShortner(urls map[string]string) *UrlShortner {
	return &UrlShortner{
		urls: urls,
	}
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
        <h2>gurl-shortner</h2>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
        <form method="post" action="/shorten">
            <input type="text" name="url" placeholder="Enter a URL">
            <input type="submit" value="Shorten">
        </form>
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

}
func (u *UrlShortner) HandleRedirectUrl(w http.ResponseWriter, r *http.Request) {}
