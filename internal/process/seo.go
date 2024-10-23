package process

import (
	"SiteSense/pkg/models"
	"fmt"
)

// SeoProcess takes SEO data and returns a formatted string
func SeoProcess(seo models.SEOMetrics) string {
	return fmt.Sprintf("H1 Tag: %s\nH1 Tag Present: %v\n", seo.MetaTitle, seo.MetaTitlePresent)
}
