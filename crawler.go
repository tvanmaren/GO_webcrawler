package main

import (
	"fmt"

	"github.com/tvanmaren/scraper"
)

func main() {
	urls := []string{"http://google.com", "http://bing.com", "http://duckduckgo.com"}
	scrapeChan := make(chan map[string][]string)
	go scraper.Scrape(urls, scrapeChan)
	linkMap := <-scrapeChan
	newLinkMap := make(map[string][]string)
	for _, links := range linkMap {
		go scraper.Scrape(links, scrapeChan)
		newMap := <-scrapeChan
		for url, newLinks := range newMap {
			newLinkMap[url] = newLinks
		}
	}
	close(scrapeChan)
	for url, link := range newLinkMap {
		fmt.Println("OVERALL", url, "CONTAINS LINKS TO", link)
	}
}
