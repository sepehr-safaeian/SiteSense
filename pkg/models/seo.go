package models

// SEOMetrics struct represents the SEO metrics of a webpage.
type SEOMetrics struct {
	// Meta Information
	MetaTitle              string   `json:"meta-title-string"`        // Meta title of the page
	MetaTitlePresent       bool     `json:"meta-title-present"`       // Indicates if the meta title tag is present
	MetaTitleLength        int      `json:"meta-title-length"`        // Length of the meta title
	MetaDescription        string   `json:"description"`              // Meta description of the page
	MetaDescriptionPresent bool     `json:"meta-description-present"` // Indicates if the meta description tag is present
	MetaDescriptionLength  int      `json:"meta-description-length"`  // Length of the meta description
	MetaKeywords           []string `json:"meta-keywords"`            // List of meta keywords
	AuthorPresent          bool
	Author                 string `json:"author"`           // Meta author tag
	LanguagePresent        bool   `json:"language-present"` // Declared language in the <html> tag
	Language               string `json:"language"`         // Declared language in the <html> tag
	ViewPortPresent        bool
	ViewPort               string `json:"viewport"` // Viewport settings from the meta tag

	// URL and Accessibility
	URL                 string `json:"url"`                   // URL of the page
	CanonicalTagPresent bool   `json:"canonical-tag-present"` // Indicates if the canonical tag is present
	RobotsTxtPresent    bool   `json:"robots-txt-present"`    // Indicates if robots.txt file is present
	XMLSitemapPresent   bool   `json:"xml-sitemap-present"`   // Indicates if XML sitemap is available

	// Link Metrics
	InternalLinksCount int `json:"internal-links-count"` // Number of internal links
	ExternalLinksCount int `json:"external-links-count"` // Number of external links

	// Image Metrics
	ImagesCount                 int      `json:"images-count"`                     // Total number of images on the page
	AltTextInImages             bool     `json:"alt-text-in-images"`               // Indicates if alt text is present for images
	HasImagesWithoutAltText     bool     `json:"has-images-without-alt-text"`      // Indicates if there are images without alt text
	CountWithoutAltTextInImages int      `json:"count-without-alt-text-in-images"` // Count of images without alt text
	ImagesWithoutAltText        []string `json:"images-without-alt-text"`          // Links to images without alt text
	AltTexts                    []string `json:"alt-texts"`                        // List of alt texts for images

	// Page Structure
	AMPEnabled     bool     `json:"amp-enabled"`     // Indicates if the page is AMP-enabled
	HasBreadcrumbs bool     `json:"has-breadcrumbs"` // Indicates if the page has breadcrumbs
	H1Tags         []string `json:"h1-tags"`         // List of H1 tags in the page
	H2Tags         []string `json:"h2-tags"`         // List of H2 tags
	ContentLength  int      `json:"content-length"`  // Total content length of the page
	ContentType    string   `json:"content-type"`    // Content type (e.g., text/html)
	PageType       string   `json:"page-type"`       // Type of page (e.g., blog post, landing page, product page)
}
