package main

import (
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
	"log"
	_ "log"
)

func main() {

	c := colly.NewCollector()

	url := "https://openweathermap.org/"

	c.OnRequest(func(request *colly.Request) {
		log.Println("Received the response. ", request.URL)

	})

	c.OnResponse(func(response *colly.Response) {
		log.Println("HTML response:", string(response.Body)) // Print the HTML body of the response
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Println("Error", err)

	})

	err := c.Visit(url)
	if err != nil {
		return
	}

}
