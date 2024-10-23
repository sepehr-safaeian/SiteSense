package collectors

import (
	"SiteSense/pkg/models"
	"log"

	"github.com/gocolly/colly"
)

// SeoCollector collects SEO metrics from the given URL and returns them
func SeoCollector(url string) models.SEOMetrics {
	var seo models.SEOMetrics

	c := colly.NewCollector()

	// H1 Collector
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		seo.MetaTitlePresent = true
		seo.MetaTitle = e.Text
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	c.Visit(url)

	return seo
}
