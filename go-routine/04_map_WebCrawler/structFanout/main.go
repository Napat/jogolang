package main

import (
	"fmt"
	"sync"
	"time"
)

type crawlData struct {
	url   string
	depth int
}

// Fetcher fetch data from url
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// CrawlCache data
type CrawlCache struct {
	Fetcher // embed field that implement interface Fetcher
	vMux    sync.Mutex
	v       map[string]string
	in      chan crawlData
}

//AskCrawlToken ask for search token
func (c *CrawlCache) AskCrawlToken(url string) bool {
	c.vMux.Lock()
	defer c.vMux.Unlock()
	if _, ok := c.v[url]; !ok {
		c.v[url] = ""
		return true
	}
	return false
}

// CrawlValue get data
func (c *CrawlCache) CrawlValue(url string) (string, bool) {
	c.vMux.Lock()
	defer c.vMux.Unlock()
	v, ok := c.v[url]
	return v, ok
}

//CrawlSave save data
func (c *CrawlCache) CrawlSave(url string, v string) {
	c.vMux.Lock()
	c.v[url] = v
	c.vMux.Unlock()
}

// Crawler uses fetcher to recursively crawl
func (c *CrawlCache) Crawler(url string, depth int) {

	if ok := c.AskCrawlToken(url); !ok {
		return
	}
	/////////////////TEST parallel////////////////////
	// force 1T = 1 sec
	time.Sleep(1000 * time.Millisecond)
	/////////////////////////////////////
	if depth <= 0 {
		return
	}
	body, urls, err := c.Fetch(url)
	c.CrawlSave(url, body)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range urls {
		c.in <- crawlData{url: u, depth: depth - 1}
	}
	return
}

func (c *CrawlCache) krawl() {
	// c.in  <-chan crawlData
	for cd := range c.in {
		c.Crawler(cd.url, cd.depth)
	}
}

func (c *CrawlCache) fanOutWorker(worker int) {
	for i := 0; i < worker; i++ {
		go c.krawl()
	}
}

// Crawl crawl url in depth
func Crawl(url string, depth int, fetcher Fetcher) *CrawlCache {
	sc := &CrawlCache{Fetcher: fetcher,
		v:  make(map[string]string),
		in: make(chan crawlData),
	}
	go sc.Crawler(url, depth)
	//sc.in<-crawlData{url:url, depth: depth}
	return sc
}

func main() {

	sc := Crawl("http://golang.org/", 4, fetcher)
	sc.fanOutWorker(4) // try change worker number

	fmt.Println("-----1-----")
	for k, v := range sc.v {
		fmt.Println(k, v)
	}
	//let's worker work for 2 sec.
	time.Sleep(2000 * time.Millisecond)

	fmt.Println("-----2-----")
	for k, v := range sc.v {
		fmt.Println(k, v)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
