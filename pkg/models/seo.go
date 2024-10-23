package models

// SEOMetrics struct
type SEOMetrics struct {
	MetaTitle               string   `json:"meta-title-string"`
	URL                     string   `json:"url"`                        // URL of the page
	MetaTitlePresent        bool     `json:"meta-title-present"`         // Meta title tag present
	MetaTitleLength         int      `json:"meta-title-length"`          // Meta title length
	MetaDescription         string   `json:"description"`                // Meta description of the page
	MetaDescriptionPresent  bool     `json:"meta-description-present"`   // Meta description tag present
	MetaDescriptionLength   int      `json:"meta-description-length"`    // Meta description length
	CanonicalTagPresent     bool     `json:"canonical-tag-present"`      // Canonical tag present
	RobotsTxtPresent        bool     `json:"robots-txt-present"`         // Robots.txt file present
	XMLSitemapPresent       bool     `json:"xml-sitemap-present"`        // XML sitemap available
	StructuredDataValid     bool     `json:"structured-data-valid"`      // Structured data (e.g., JSON-LD, Microdata)
	HreflangTagsPresent     bool     `json:"hreflang-tags-present"`      // Hreflang tags for multilingual websites
	InternalLinksCount      int      `json:"internal-links-count"`       // Number of internal links
	ExternalLinksCount      int      `json:"external-links-count"`       // Number of external links
	AltTextInImages         bool     `json:"alt-text-in-images"`         // Alt text for images
	MobileFriendly          bool     `json:"mobile-friendly"`            // Mobile-friendliness
	PageSpeedScore          int      `json:"page-speed-score"`           // Page speed score (Google Lighthouse)
	SocialSharingTags       bool     `json:"social-sharing-tags"`        // Open Graph, Twitter Cards for social media sharing
	NoIndexNoFollowPresent  bool     `json:"no-index-no-follow-present"` // Noindex, nofollow tags for pages
	BacklinkProfileStrength int      `json:"backlink-profile-strength"`  // Strength of backlink profile
	MetaKeywords            []string `json:"meta-keywords"`              // List of meta keywords
	Language                string   `json:"language"`                   // Declared language in the <html> tag
	ViewPort                string   `json:"viewport"`                   // Viewport settings from meta tag
	Author                  string   `json:"author"`                     // Meta author tag
	PublishDate             string   `json:"publish-date"`               // Publish date (if applicable)
	LastModifiedDate        string   `json:"last-modified-date"`         // Last modified date
	ContentLength           int      `json:"content-length"`             // Total content length of the page
	ContentType             string   `json:"content-type"`               // Content type (e.g., text/html)
	PageType                string   `json:"page-type"`                  // Type of page (e.g., blog post, landing page, product page)
	AMPEnabled              bool     `json:"amp-enabled"`                // Is the page AMP-enabled?
	HasBreadcrumbs          bool     `json:"has-breadcrumbs"`            // Does the page have breadcrumbs?
	H1Tags                  []string `json:"h1-tags"`                    // List of H1 tags in the page
	H2Tags                  []string `json:"h2-tags"`                    // List of H2 tags
	ImagesCount             int      `json:"images-count"`               // Total number of images on the page
}
