// https://tour.golang.org/concurrency/10

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Entry struct{
	body string
	err error
}

// Cache is a lockable result store
type Cache struct {
	m   map[string]Entry
	log []string
	mux sync.Mutex
}

// Log adds to the log and optionally prints if verbose is true
func (c *Cache) Log(s string, verb bool) {
	c.log = append(c.log, s)
	if verb {
		fmt.Println(s)
	}
}

var results = Cache{m: make(map[string]Entry)}


func crawlConc(url string, depth int, fetcher Fetcher, verb bool, crawled chan bool) {
	Crawl(url, depth, fetcher, verb)
	crawled <- true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, verb bool) (mess string) {
	if depth <= 0 {
		return
	}
	// lock before reading from results chache
	results.mux.Lock()
	_, exists := results.m[url]
	if !exists {
		// Let other goroutines know this url is already being handled
		results.m[url] = Entry{"fetching", nil}
		// Log and optionally print
		results.Log(fmt.Sprintf("fetching: %s", url), verb)
	}

	// unlock results cache for other goroutines to use
	results.mux.Unlock()
	if exists {
		results.Log(
			fmt.Sprintf("Nothing to do: %s has already been visited", url), verb)
		return
	}

	body, urls, err := fetcher.Fetch(url)

	results.mux.Lock()
	results.m[url] = Entry{body, err}
	results.mux.Unlock()

	if err != nil {
		results.Log(fmt.Sprintf("%v", err), verb)
		return
	}
	results.Log(fmt.Sprintf("found: %s %q", url, body), verb)

	crawled := make(chan bool, len(url))
	for _, u := range urls {
		go crawlConc(u, depth-1, fetcher, verb, crawled)
	}
	for range urls {
		<-crawled
	}

	return fmt.Sprintf("Verbose: %v\nTo view log, see 'results.log'.\nTo view crawled urls see 'results.m'.", verb)
}

func main() {
	fmt.Println("Running Crawl (verbose: false)...")
	fmt.Println(Crawl("https://golang.org/", 4, fetcher, false))
	fmt.Println("Printing Crawl results...")
	i := 0
	for u, r := range(results.m) {
		if r.err != nil {
			i++
			defer fmt.Printf("Error! %v\n", r.err)
			continue
		}
		fmt.Printf("Success! found: %s with body: \"%s\"\n", u, r.body)
	}
	fmt.Printf("Errors: %d\n", i)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
