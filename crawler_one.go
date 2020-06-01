package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var target string
var urls = make(map[string]string)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please input URL")
		return
	}

	target = os.Args[1]
	RunCrawler(target)
}

func RunCrawler(uri string) {
	if uri == "" {
		return
	}
	if urls[uri] == "" {
		urls[uri] = uri
	} else {
		return
	}

	fmt.Println("Fetch", uri)
	response, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	doc.Find("a").Each(func(i int, q *goquery.Selection) {
		attr, exists := q.Attr("href")
		if exists {
			nextLink := TrimUrl(attr)
			RunCrawler(nextLink)
		}
	})
}

func TrimUrl(uri string) string {
	uri = strings.TrimSuffix(uri, "/")
	validUrl, err := url.Parse(uri)
	if err != nil {
		return ""
	}

	targetUri, _ := url.Parse(target)
	if strings.Contains(validUrl.String(), targetUri.Host) {
		return uri
	}

	return ""
}
