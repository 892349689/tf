package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"net/http"
	"os"
)
func catalog(url string) *html.Node {
	resp,err := http.Get(url)
	if err !=nil {
		os.Exit(0)
	}
	defer resp.Body.Close()
	body, err := htmlquery.Parse(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		os.Exit(0)
	}
	return body
}

func request(url string) {
	content := catalog(url)
	nodes := htmlquery.Find(content,"//div[@class='volume-list']//div//ul//li/a")
	for  {
		rell := htmlquery.SelectAttr(nodes[k], "href")
		if *section != "" {
			if rell == *section {
				fetch(rell)
				*section = ""
			} else {
				k += 1
			}
		} else {
			fetch(rell)
		}
	}
}
