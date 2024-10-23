package collectors

import (
	"SiteSense/pkg/models"
	"log"

	"github.com/gocolly/colly"
)

// SeoCollector collects SEO metrics from the given URL and returns them
func SeoCollector(url string) (models.SEOMetrics, string) {
	var seo models.SEOMetrics

	c := colly.NewCollector()

	// Collect the URL
	seo.URL = url // Assuming you have added a URL field to the SEOMetrics struct

	// Set up the collector to find h1 and meta description
	collectH1(c, &seo)
	collectMetaDescription(c, &seo)

	// Error handling
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Start visiting the URL
	err := c.Visit(url)
	if err != nil {
		log.Println("Error visiting URL:", err)
	}

	return seo, url // Return both SEO data and the URL
}

// collectH1 extracts the H1 tag from the page
func collectH1(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		seo.MetaTitlePresent = true
		seo.MetaTitle = e.Text
		seo.MetaTitleLength = len(seo.MetaTitle)
	})
}

// collectMetaDescription extracts the meta description from the page
func collectMetaDescription(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		seo.MetaDescriptionPresent = true
		seo.MetaDescription = e.Attr("content")
		seo.MetaDescriptionLength = len(seo.MetaDescription)
	})
}
