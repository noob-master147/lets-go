package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(2 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		fmt.Println("Error: ", err)
		c <- link
		return
	}
	fmt.Println(link, "is Up!")
	c <- link
}
