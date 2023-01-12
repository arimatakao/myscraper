package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Product struct {
	Name         string
	Price        string
	Rating       string
	Availability string
	Link         string
	Image        string
}

func NewProduct(name, price, rating, availability, link, image string) *Product {
	return &Product{
		Name:         name,
		Price:        price,
		Rating:       rating,
		Availability: availability,
		Link:         link,
		Image:        image,
	}
}

func main() {
	c := colly.NewCollector()
	// products := make([]Product, 0)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	// Find products
	// c.OnHTML("ul.catalog-grid", func(e *colly.HTMLElement) {
	c.OnHTML("ul.catalog-grid", func(e *colly.HTMLElement) {
		fmt.Println("Start scarping")
		// e.ForEach("div.goods-tile__inner",
		e.ForEach("li.catalog-grid__cell",
			func(i int, h *colly.HTMLElement) {
				// item := Product{}
				// item.Name = h.Text
				// item.Price = h.Attr("goods-tile__price-value")
				// item.Rating = h.Attr("goods-tile__reviews-link")
				// fmt.Println(item.Name+" | ", item.Price+" | ", item.Availability)
				p := &Product{}
				p.Name = h.ChildText("span.goods-tile__title")
				p.Price = h.ChildText("span.goods-tile__price-value")
				p.Rating = h.ChildText("div.goods-tile__stars")
				p.Availability = h.ChildText("div.goods-tile__availability")
				fmt.Println(p.Name+" | ", p.Price+" | ", p.Rating+" | ", p.Availability)
			})
	})
	// example link
	c.Visit("https://rozetka.com.ua/ua/notebooks/c80004/")
}
