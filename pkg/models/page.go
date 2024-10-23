package models

// Additional page-level metrics structs can be referenced here
type Page struct {
	Performance   PerformanceMetrics   `json:"performance"`   // Performance metrics of the page
	SEO           SEOMetrics           `json:"seo"`           // SEO metrics of the page
	Accessibility AccessibilityMetrics `json:"accessibility"` // Accessibility metrics of the page
	Security      SecurityMetrics      `json:"security"`      // Security metrics of the page
	Usability     UsabilityMetrics     `json:"usability"`     // Usability metrics of the page
}
