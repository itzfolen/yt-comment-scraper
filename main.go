package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
)

func main() {
	args := os.Args
	url := args[1]
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})
	collector.Visit(url)
}
