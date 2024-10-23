package models

// SEOMetrics struct
type SEOMetrics struct {
	MetaTitle               string `json:"meta-title-string"`
	MetaTitlePresent        bool   `json:"meta-title-present"`         // Meta title tag present
	MetaTitleLength         int    `json:"meta-title-length"`          // Meta title length
	MetaDescriptionPresent  bool   `json:"meta-description-present"`   // Meta description tag present
	MetaDescriptionLength   int    `json:"meta-description-length"`    // Meta description length
	CanonicalTagPresent     bool   `json:"canonical-tag-present"`      // Canonical tag present
	RobotsTxtPresent        bool   `json:"robots-txt-present"`         // Robots.txt file present
	XMLSitemapPresent       bool   `json:"xml-sitemap-present"`        // XML sitemap available
	StructuredDataValid     bool   `json:"structured-data-valid"`      // Structured data (e.g., JSON-LD, Microdata)
	HreflangTagsPresent     bool   `json:"hreflang-tags-present"`      // Hreflang tags for multilingual websites
	InternalLinksCount      int    `json:"internal-links-count"`       // Number of internal links
	ExternalLinksCount      int    `json:"external-links-count"`       // Number of external links
	AltTextInImages         bool   `json:"alt-text-in-images"`         // Alt text for images
	MobileFriendly          bool   `json:"mobile-friendly"`            // Mobile-friendliness
	PageSpeedScore          int    `json:"page-speed-score"`           // Page speed score (Google Lighthouse)
	SocialSharingTags       bool   `json:"social-sharing-tags"`        // Open Graph, Twitter Cards for social media sharing
	NoIndexNoFollowPresent  bool   `json:"no-index-no-follow-present"` // Noindex, nofollow tags for pages
	BacklinkProfileStrength int    `json:"backlink-profile-strength"`  // Strength of backlink profile
}
