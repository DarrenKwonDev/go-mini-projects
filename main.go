package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type result struct {
	url    string
	status string
}

func main() {

	c := make(chan result)

	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://cineps.net/"}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

func hitURL(url string, c chan result) {
	fmt.Println("checking", url)

	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		c <- result{url: url, status: "FALED"}
	}
	c <- result{url: url, status: status}
}
