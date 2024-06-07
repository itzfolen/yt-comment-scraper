package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
)

type Dictionary map[string]string

type RecipeSpecs struct {
	difficulty, prepTime, cookingTime, servingSize, priceTier string
}

type Recipe struct {
	url, name      string
	ingredients    []Dictionary
	specifications RecipeSpecs
}

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
