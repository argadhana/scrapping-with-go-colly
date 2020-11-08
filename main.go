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
			var title, con, rating, image, tipe, link string
			title = e.ChildText("div.tt")

			if title == "" {
				return
			}

			con = e.ChildText("div.epxs")
			if con == "" {
				return
			}

			rating = e.ChildText("i")
			if rating == "" {
				return
			}

			image = e.ChildAttr("img", "src")
			if image == "" {
				return
			}

			tipe = e.ChildText("span.type")
			if tipe == "" {
				return
			}

			link = e.ChildAttr("a", "href")
			if link == "" {
				return
			}

			fmt.Printf("Judul Komik: %s \nCH: %s \nRating: %s \nImage: %s \nTipe: %s \nLink: %s \n\n", title, con, rating, image, tipe, link)

		})
	})

	scrapp.Visit("https://komikcast.com/komik/")
}
