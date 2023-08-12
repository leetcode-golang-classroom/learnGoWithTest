package concurrency

import "time"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		// results[url] = wc(url)
		go func(url string) {
			// results[url] = wc(url)
			resultChannel <- result{url, wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	time.Sleep(2 * time.Second)
	return results
}
