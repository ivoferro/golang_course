package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.com",
		"http://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	<-c
	// }

	// for range links {
	// 	<-c
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	for link := range c {
		go func(innerLink string) {
			time.Sleep(5 * time.Second)
			checkLink(innerLink, c)
		}(link)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "seems to be down.")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
