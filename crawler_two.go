package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

var target string
var wg sync.WaitGroup

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please input URL")
		return
	}

	target = os.Args[1]
	queue := make(chan string)
	filterQueue := make(chan string)

	wg.Add(1)
	go func() {
		queue <- target
	}()

	go func() {
		var urls = make(map[string]bool)
		for uri := range queue {
			if !urls[uri] {
				urls[uri] = true
				filterQueue <- uri
			} else { wg.Done() }
		}
	}()

	for i := 0; i < 5; i++ {
		go func() {
			for uri := range filterQueue {
				RunCrawler(uri, queue)
				wg.Done()
			}
		}()
	}

	wg.Wait()
}

func RunCrawler(uri string, ch chan string) {
	if uri == "" {
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
			wg.Add(1)
			go func() {
				ch <- nextLink
			}()
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
