package cmd

import (
	"fmt"
	"time"

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
		fmt.Printf(config.Red+config.Bold+"(%s)\n"+config.Reset+"All packages are available on: "+config.Bold+"%s\n"+config.Reset, time.Now().Format("2006-01-02 15:04:05"), r.URL)
	})

	c.Visit(packagesurl)
}
