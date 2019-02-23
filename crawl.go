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

func query() {

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

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find(".rank").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Post %s, %s\n", link, linkText)
	})
}
