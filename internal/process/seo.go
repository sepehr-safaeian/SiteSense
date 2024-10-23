package process

import (
	"SiteSense/pkg/models"
	"strconv"
	"strings"
)

// SeoProcess takes SEO data and returns a formatted string
func SeoProcess(seo models.SEOMetrics) string {
	var builder strings.Builder

	// Append the SEO metrics to the builder
	builder.WriteString("SEO Metrics:\n")
	appendMetric(&builder, "\nURL", seo.URL)
	appendMetric(&builder, "\nTitle(h1) Tag Present", boolToStr(seo.MetaTitlePresent))
	appendMetric(&builder, "Title(h1) Tag Text", seo.MetaTitle)
	appendMetric(&builder, "Title(h1) Tag Length", strconv.Itoa(seo.MetaTitleLength))

	appendMetric(&builder, "\nMeta Description Present", boolToStr(seo.MetaDescriptionPresent))
	appendMetric(&builder, "Meta Description", seo.MetaDescription)
	appendMetric(&builder, "Meta Description Length", strconv.Itoa(seo.MetaDescriptionLength))

	return builder.String()
}

// appendMetric appends a key-value pair to the StringBuilder
func appendMetric(builder *strings.Builder, key, value string) {
	builder.WriteString(key + ": " + value + "\n")
}

// boolToStr converts a boolean to a string
func boolToStr(b bool) string {
	if b {
		return "✅"
	}
	return "❌"
}
