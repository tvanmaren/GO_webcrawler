package main

import (
	"fmt"
	// "os"

	"github.com/tvanmaren/scraper"
)

func main() {
	// depth, _ := strconv.Atoi(os.Args[1])
	depth := 2
	urls := []string{"http://google.com", "http://bing.com", "http://duckduckgo.com"}
	scrapeChan := make(chan map[string][]string)
	fmt.Println("SCRAPING", urls)
	linkMap := make(map[string][]string)
	go scraper.Scrape(urls, scrapeChan)
	for url, links := range <-scrapeChan {
		fmt.Println("HEARD SCRAPE OF", url)
		linkMap[url] = links
	}
	crawlChan := make(chan map[string][]string)
	fmt.Println("CRAWLING", linkMap)
	go crawl(depth, linkMap, crawlChan)
	for url, links := range <-crawlChan {
		fmt.Println("OVERALL", url, "CONTAINS LINKS TO", links)
		linkMap[url] = links
	}
}

func crawl(depth int, linkMap map[string][]string, scrapeChan chan map[string][]string) {
	if depth == 0 {
		scrapeChan <- linkMap
		close(scrapeChan)
		return
	}
	newLinkMap := make(map[string][]string)
	for _, links := range linkMap {
		scrapeChan2 := make(chan map[string][]string)
		fmt.Println("SENDING SCRAPE FOR", links)
		go scraper.Scrape(links, scrapeChan2)
		for url, newLinks := range <-scrapeChan2 {
			fmt.Println("HEARD BACK FROM SCRAPE OF", url)
			newLinkMap[url] = newLinks
		}
		crawlChan := make(chan map[string][]string)
		go crawl(depth-1, newLinkMap, crawlChan)
		for url, newLinks := range <-crawlChan {
			fmt.Println("SUCCESSFULLY CRAWLED", url)
			linkMap[url] = newLinks
		}
	}
	scrapeChan <- linkMap
}
