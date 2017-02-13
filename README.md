# Basic WebCrawler in GO
___
## Operations
This webcrawler uses Golang's github integration to pull in another GO repository I made for a [URL Webscraper](github.com/tvanmaren/scraper). With it, we grab the HTML from a given URL, parse it for links to other pages, and pull together a map of what links can be found on what pages, much as basic search engine would do to create its database of information.
This webcrawler is currently hardcoded to operate at two full layers of depth. (see diagram below)
```
layer 0:                                                bing.com
                                                          |
                                                        /   \
layer 1:                                  bing.com/news     bing.com/images
                                                  |
                                                /   \
layer 2:                                npr.org      bing.com/images
                                          |
                                        /   \
layer 2's links:            help.npr.org    npr.org/contact
```
___
## How to Use
As this repository is publicly available, if you already have Golang installed on your machine, you can just type `go get https://github.com/tvanmaren/GO_webcrawler` and you should be able to run it like a charm.
___
## Warnings
This crawler currently uses three seed pages:

1. Duckduckgo.com
  * contains no links to other pages
  * completes map immediately
2. Google.com
  * contains straightforward links to google mail, images, and apps
  * successfully finishes two layer map after exhausting a few link-intensive pages
3. Bing.com
  * contains myriad direct links to current news articles, images, microsoft services, and more
  * I've never gotten it to successfully finish

Because bing.com contains such an absurd number of direct links right off, the project of a two-layers-deep map becomes exponentially more intensive. I have left it for hours without it finishing. So if you're going to run this webcrawler, just expect to manually cut it off at some point, otherwise it will keep bouncing back and forth between news articles, other news articles, services that you aren't logged into, info pages, etc,
