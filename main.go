package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yterajima/go-sitemap"
)

var anyURLError = false

func checkURL(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		anyURLError = true
		fmt.Println("============")
		fmt.Println(url)
		fmt.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(":: Error! Not sitemap url supplied.")
		fmt.Println(":: Usage: checksitemap https://example.com/sitemap.xml")

		return
	}

	sitemapURL := os.Args[1]

	sitemap, err := sitemap.Get(sitemapURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("> Sitemap seems to be present.")
	fmt.Println("> Checking links...")

	for _, URL := range sitemap.URL {
		checkURL(URL.Loc)
	}

	if anyURLError {
		fmt.Println("============")
	}
	fmt.Println("> Check complete!")
}
