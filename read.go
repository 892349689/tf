package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/eiannone/keyboard"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"os/exec"
	"time"
)


func read(url string) *html.Node{
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error", err)
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

func fetch(url string)  {
	content := read(url)
	nodes := htmlquery.Find(content, "//div[@class='content']//p")
	n := 0
	M:
	for {
		defer func() {
			err := recover()
			if err != nil {
				k += 1
				return
			}
		}()
		line := htmlquery.FindOne(nodes[n], "//p/text()").Data
		cmd := exec.Command("cmd.exe", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("\r%s\r", line)
		if *num != 0{
			time.Sleep(time.Duration(*num) * time.Second)
			n += 1
			continue
		}else {
			for {
				keysEvents, err := keyboard.GetKeys(10)
				if err != nil {
					panic(err)
				}
				defer func() {
					_ = keyboard.Close()
				}()
				event := <-keysEvents
				if event.Key == keyboard.KeyArrowUp {
					if n > 0 {
						n -= 1
					}
					break
				}
				if event.Key == keyboard.KeyArrowDown {
					n += 1
					break
				}
				if event.Key == keyboard.KeyArrowLeft {
					if k > 0 {
						k -= 1
					}
					break M
				}
				if event.Key == keyboard.KeyArrowRight {
					k += 1
					break M
				}
			}
		}
	}
}
