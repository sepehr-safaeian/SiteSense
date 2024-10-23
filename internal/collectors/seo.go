package collectors

import (
	"SiteSense/pkg/models"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

// SeoCollector collects SEO metrics from the given URL and returns them
func SeoCollector(url string) (models.SEOMetrics, string) {
	var seo models.SEOMetrics
	seo.URL = url // Collect the URL

	c := colly.NewCollector()

	// Set up collectors
	setupCollectors(c, &seo)

	// Error handling
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Start visiting the URL
	if err := c.Visit(url); err != nil {
		log.Println("Error visiting URL:", err)
	}

	return seo, url // Return both SEO data and the URL
}

// setupCollectors initializes all collectors
func setupCollectors(c *colly.Collector, seo *models.SEOMetrics) {
	collectH1(c, seo)
	collectMetaDescription(c, seo)
	collectCanonicalTag(c, seo)
	collectSitemapXML(c, seo)
	collectInternalLinks(c, seo)
	collectExternalLinks(c, seo)
	collectImages(c, seo)
	collectAuthor(c, seo)
	collectLanguage(c, seo)
	collectViewPort(c, seo)
	collectRobotsTxtPresent(c, seo)
}

// collectH1 extracts the H1 tag from the page
func collectH1(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		seo.MetaTitlePresent = true
		seo.MetaTitle = e.Text
		seo.MetaTitleLength = len(seo.MetaTitle)
	})
}

// collectCanonicalTag checks for the canonical link
func collectCanonicalTag(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML(`link[rel="canonical"]`, func(e *colly.HTMLElement) {
		seo.CanonicalTagPresent = true
	})
}

// collectSitemapXML collects URLs from the XML sitemap
func collectSitemapXML(c *colly.Collector, seo *models.SEOMetrics) {
	sitemapURL := seo.URL + "/sitemap_index.xml"

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			seo.XMLSitemapPresent = true
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		seo.XMLSitemapPresent = false
	})

	if err := c.Visit(sitemapURL); err != nil {
		seo.XMLSitemapPresent = false
	}
}

// collectMetaDescription extracts the meta description from the page
func collectMetaDescription(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		seo.MetaDescriptionPresent = true
		seo.MetaDescription = e.Attr("content")
		seo.MetaDescriptionLength = len(seo.MetaDescription)
	})
}

// collectInternalLinks extracts and counts internal links from the page
func collectInternalLinks(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if isInternalLink(e.Attr("href"), seo.URL) {
			seo.InternalLinksCount++
		}
	})
}

// isInternalLink checks if a link is internal
func isInternalLink(link, baseURL string) bool {
	if link == "" {
		return false
	}

	if strings.HasPrefix(link, "http") {
		return strings.HasPrefix(link, baseURL)
	}

	return true
}

// collectExternalLinks extracts and counts external links from the page
func collectExternalLinks(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if isExternalLink(e.Attr("href"), seo.URL) {
			seo.ExternalLinksCount++
		}
	})
}

// isExternalLink checks if a link is external
func isExternalLink(link, baseURL string) bool {
	if link == "" {
		return false
	}

	return strings.HasPrefix(link, "http") && !strings.HasPrefix(link, baseURL)
}

// collectImages extracts images, counts them, checks for alt text, and stores links without alt text
func collectImages(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		seo.ImagesCount++
		altText := e.Attr("alt")

		if altText == "" {
			seo.CountWithoutAltTextInImages++
			seo.ImagesWithoutAltText = append(seo.ImagesWithoutAltText, e.Attr("src"))
		} else {
			seo.AltTexts = append(seo.AltTexts, altText)
		}
	})
}

// collectAuthor collects author meta information
func collectAuthor(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML(`meta[name="author"]`, func(e *colly.HTMLElement) {
		seo.AuthorPresent = true
		seo.Author = e.Attr("content")
	})

	c.OnScraped(func(r *colly.Response) {
		if !seo.AuthorPresent {
			log.Println("No author meta tag found for:", r.Request.URL)
		}
	})
}

// collectLanguage collects language information from the HTML tag
func collectLanguage(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML("html", func(e *colly.HTMLElement) {
		if lang := e.Attr("lang"); lang != "" {
			seo.LanguagePresent = true
			seo.Language = lang
		}
	})

	c.OnScraped(func(r *colly.Response) {
		if !seo.LanguagePresent {
			log.Println("No lang attribute found for:", r.Request.URL)
		}
	})
}

// collectViewPort collects ViewPort meta information
func collectViewPort(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML(`meta[name="viewport"]`, func(e *colly.HTMLElement) {
		seo.ViewPortPresent = true
		seo.ViewPort = e.Attr("content")
		log.Println("Viewport collected:", seo.ViewPort) // Log the collected viewport
	})

	c.OnScraped(func(r *colly.Response) {
		if !seo.ViewPortPresent {
			log.Println("No viewport meta tag found for:", r.Request.URL)
		} else {
			log.Println("Viewport tag is present for:", r.Request.URL)
		}
	})
}

// collectRobotsTxtPresent collects RobotsTxtPresent meta information
func collectRobotsTxtPresent(c *colly.Collector, seo *models.SEOMetrics) {
	c.OnHTML(`meta[name="robots"]`, func(e *colly.HTMLElement) {
		seo.RobotsTxtPresent = true
		log.Println("Viewport collected:", seo.RobotsTxtPresent) // Log the collected RobotsTxtPresent
	})

	c.OnScraped(func(r *colly.Response) {
		if !seo.ViewPortPresent {
			log.Println("No viewport meta tag found for:", r.Request.URL)
		} else {
			log.Println("Viewport tag is present for:", r.Request.URL)
		}
	})
}
