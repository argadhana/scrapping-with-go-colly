package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	scrapp := colly.NewCollector()

	scrapp.OnRequest(func(r *colly.Request) {
		fmt.Println("Menelusuri", r.URL)
	})
	scrapp.OnHTML("div.listo", func(e *colly.HTMLElement) {
		e.ForEach("div.bs", func(_ int, e *colly.HTMLElement) {
			var title, con, rating string

			title = e.ChildText("div.tt")

			if title == "" {
				return
			}

			con = e.ChildText("div.epxs")
			if con == "" {
				return
			}

			rating = e.ChildText("div.rating")
			if rating == "" {
				return
			}

			fmt.Printf("Judul Komik: %s \nCH: %s \nRating: %s \n", title, con, rating)

		})
	})

	scrapp.Visit("https://komikcast.com/komik/")
}
