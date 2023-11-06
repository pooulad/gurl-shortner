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

}
func (u *UrlShortner) HandleRedirectUrl(w http.ResponseWriter, r *http.Request) {}
