package cmd

import (
	"fmt"

	"github.com/gocolly/colly"
	config "github.com/rendick/pem/settings"
)

var (
	// website
	packagesurl = "https://pem-packages.vercel.app/"
)

func Scrapper() {
	c := colly.NewCollector()

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		fmt.Println(config.Red + config.Bold + e.Text + config.Reset)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("All packages are available on:", r.URL)
	})

	c.Visit(packagesurl)
}
