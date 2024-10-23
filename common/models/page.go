package models

// PageInfo struct
type PageInfo struct {
	URL              string   `json:"url"`                // URL of the page
	Title            string   `json:"title"`              // Title of the page
	Description      string   `json:"description"`        // Meta description of the page
	MetaKeywords     []string `json:"meta-keywords"`      // List of meta keywords
	CanonicalURL     string   `json:"canonical-url"`      // Canonical URL
	Language         string   `json:"language"`           // Declared language in the <html> tag
	ViewPort         string   `json:"viewport"`           // Viewport settings from meta tag
	Author           string   `json:"author"`             // Meta author tag
	PublishDate      string   `json:"publish-date"`       // Publish date (if applicable)
	LastModifiedDate string   `json:"last-modified-date"` // Last modified date
	ContentLength    int      `json:"content-length"`     // Total content length of the page
	ContentType      string   `json:"content-type"`       // Content type (e.g., text/html)
	PageType         string   `json:"page-type"`          // Type of page (e.g., blog post, landing page, product page)
	AMPEnabled       bool     `json:"amp-enabled"`        // Is the page AMP-enabled?
	HasBreadcrumbs   bool     `json:"has-breadcrumbs"`    // Does the page have breadcrumbs?
	LinksInternal    int      `json:"links-internal"`     // Number of internal links
	LinksExternal    int      `json:"links-external"`     // Number of external links
	H1Tags           []string `json:"h1-tags"`            // List of H1 tags in the page
	H2Tags           []string `json:"h2-tags"`            // List of H2 tags
	ImagesCount      int      `json:"images-count"`       // Total number of images on the page
	ScriptCount      int      `json:"script-count"`       // Total number of scripts loaded
	StylesheetCount  int      `json:"stylesheet-count"`   // Total number of stylesheets loaded
	PageLoadTime     float32  `json:"page-load-time"`     // Page load time (in seconds)
}

// Additional page-level metrics structs can be referenced here
type Page struct {
	Info          PageInfo             `json:"info"`          // General information about the page
	Performance   PerformanceMetrics   `json:"performance"`   // Performance metrics of the page
	SEO           SEOMetrics           `json:"seo"`           // SEO metrics of the page
	Accessibility AccessibilityMetrics `json:"accessibility"` // Accessibility metrics of the page
	Security      SecurityMetrics      `json:"security"`      // Security metrics of the page
	Usability     UsabilityMetrics     `json:"usability"`     // Usability metrics of the page
}
