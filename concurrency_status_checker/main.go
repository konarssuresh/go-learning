package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(url string, c chan string) {

	_, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occured while fetching ", url)
		c <- url
		return
	}
	fmt.Println("link is up", url)
	c <- url
}
func main() {
	urls := []string{"https://www.facebook.com", "https://www.google.com", "https://www.twitter.com", "https://www.instagram.com"}
	c := make(chan string)
	for _, url := range urls {
		go checkLink(url, c)

	}
	for l := range c {

		go func(url string) {
			time.Sleep(5 * time.Second)
			checkLink(url, c)
		}(l)
	}
}
