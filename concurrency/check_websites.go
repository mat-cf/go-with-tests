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
		// colocar go na frente cria uma goroutine (concurrency)
		go func() { // função anonima
			// Send statement
			resultChannel <- result{url, wc(url)}
		}()
	}

	// escrevendo de uma maneira mais otimizada
	for range urls {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	time.Sleep(2 * time.Second)
	return results
}
