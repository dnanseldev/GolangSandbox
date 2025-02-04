package concurrency_examples

import (
	"io"
	"net/http"
	"regexp"
)

func Title(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			result, _ := http.Get(url)
			html, _ := io.ReadAll(result.Body)

			r, _ := regexp.Compile(`<title>(.*?)</title>`)
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return ch
}
