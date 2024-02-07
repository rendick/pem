package cmd

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	config "github.com/rendick/pem/settings"
)

var (
	PackagesURL = "https://pem-packages.vercel.app/"
)

func Scrapper() {
	c := colly.NewCollector()

	c.OnHTML("ul#myMenu", func(e *colly.HTMLElement) {
		fmt.Println(config.Red + config.Bold + e.Text + config.Reset)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf(config.Red+config.Bold+config.Reset+"All packages are available on: "+config.Bold+"%s"+config.Reset, r.URL)
	})

	c.Visit(PackagesURL)
}
