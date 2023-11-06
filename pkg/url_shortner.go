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

}
func (u *UrlShortner) HandleRedirectUrl(w http.ResponseWriter, r *http.Request) {}
