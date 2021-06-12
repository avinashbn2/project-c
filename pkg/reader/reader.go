package reader

import (
	"log"
	"net/http"
	"net/url"

	readability "github.com/go-shiori/go-readability"
)

var (
	urls = []string{
		"https://www.awesome-react-native.com/",
		"https://github.com/inancgumus/learngo",
	}
)

type FetchItem struct {
	URL      string
	Title    string
	Author   string
	Length   int
	Excerpt  string
	SiteName string
	Image    string
}

func Fetch(urlString string) (*FetchItem, error) {
	u := urlString
	uri, err := url.Parse(u)
	if err != nil {
		log.Printf("failed to parse url %s: %v\n", u, err)
		return nil, err
	}
	resp, err := http.Get(u)
	if err != nil {
		log.Printf("failed to download this error %s: %v\n", u, err)
		return nil, err
	}
	defer resp.Body.Close()

	article, err := readability.FromReader(resp.Body, uri)
	if err != nil {
		log.Printf("failed to pars %s: %v\n", u, err)
		return nil, err
	}

	return &FetchItem{URL: urlString, Title: article.Title, Author: article.Byline, Length: article.Length, SiteName: article.SiteName, Image: article.Image, Excerpt: article.Excerpt}, nil
}
