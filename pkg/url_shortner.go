package pkg

type UrlShortner struct {
	urls map[string]string
}

func NewUrlShortner(urls map[string]string) *UrlShortner {
	return &UrlShortner{
		urls: urls,
	}
}

