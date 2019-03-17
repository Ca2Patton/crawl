package crawl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	// TOP500 - The top 500 most important websites
	TOP500 = "https://moz.com/top500"
	BLOG   = "https://jonathanmh.com/"
)

func query() []string {

	client := http.Client{}
	//FOO ***
	//req, err := http.NewRequest("GET", BLOG, nil)
	req, err := http.NewRequest("GET", TOP500, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//doc, err := goquery.NewDocument("http://jonathanmh.com")
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	// top 500 urls appened to...
	var urls []string

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find(".url").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Post %s, %s\n", link, linkText)
		urls = append(urls, link)
	})

	return urls
}

func fetch() {
	for _, url := range query() {
		client := http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("title").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			title := s.Find("title").Text()
			fmt.Printf("Title: %s\n", title)
		})
	}
}
