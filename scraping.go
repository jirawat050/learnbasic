package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	
)


func ExampleScrape() {
	// Request the HTML page.

	res, err := http.Get("https://www.edtguide.com/place/%E0%B8%AB%E0%B8%A1%E0%B8%A7%E0%B8%94%E0%B9%82%E0%B8%A3%E0%B8%87%E0%B8%9E%E0%B8%A2%E0%B8%B2%E0%B8%9A%E0%B8%B2%E0%B8%A5/%E0%B8%AA%E0%B8%A1%E0%B8%B8%E0%B8%97%E0%B8%A3%E0%B8%9B%E0%B8%A3%E0%B8%B2%E0%B8%81%E0%B8%B2%E0%B8%A3/view/P0")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".archive-list article a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("p").Text()
		href, exists := s.Attr("href")
		if exists {
			//	fmt.Println(href)
		}
		title := s.Find("h2").Text()
		log.Printf("Review %d: %s  - %s - %s\n", i, band, title, href)

	})

}

func main() {
	ExampleScrape()

}
