package main

import (
	"fmt"

	"github.com/tvanmaren/scraper"
)

func main() {
	urls := []string{"http://google.com", "http://bing.com", "http://duckduckgo.com"}
	linkMap := scraper.Scrape(urls)
	for url, links := range linkMap {
		fmt.Println(url, links)
	}
}
