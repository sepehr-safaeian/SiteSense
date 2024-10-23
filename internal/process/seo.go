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
	appendMetric(&builder, "\nMeta Title Tag Present", boolToStr(seo.MetaTitlePresent))
	appendMetric(&builder, "Meta Title Tag Text", seo.MetaTitle)
	appendMetric(&builder, "Meta Title Tag Length", strconv.Itoa(seo.MetaTitleLength))

	appendMetric(&builder, "\nMeta Description Present", boolToStr(seo.MetaDescriptionPresent))
	appendMetric(&builder, "Meta Description", seo.MetaDescription)
	appendMetric(&builder, "Meta Description Length", strconv.Itoa(seo.MetaDescriptionLength))

	appendMetric(&builder, "\nCanonical Tag Present", boolToStr(seo.CanonicalTagPresent))
	appendMetric(&builder, "\nSitemap XML Present", boolToStr(seo.XMLSitemapPresent))
	appendMetric(&builder, "\nInternal Links Count", strconv.Itoa(seo.InternalLinksCount))
	appendMetric(&builder, "\nExternal Links Count", strconv.Itoa(seo.ExternalLinksCount))

	// Append image metrics
	appendMetric(&builder, "\nImages Count", strconv.Itoa(seo.ImagesCount))
	appendMetric(&builder, "\nImages with Alt Text Present", boolToStr(seo.AltTextInImages))

	// Append count of images without alt text
	appendMetric(&builder, "\nCount of Images without Alt Text", strconv.Itoa(seo.CountWithoutAltTextInImages))

	appendMetric(&builder, "\nAuthorPresent", boolToStr(seo.AuthorPresent))
	if seo.AuthorPresent {
		appendMetric(&builder, "\nAuthor", (seo.Author))
	}

	appendMetric(&builder, "\nLanguagePresent", boolToStr(seo.LanguagePresent))
	if seo.LanguagePresent {
		appendMetric(&builder, "\nLanguage", (seo.Language))
	}

	appendMetric(&builder, "\nViewPort", boolToStr(seo.ViewPortPresent))
	if seo.ViewPortPresent {
		appendMetric(&builder, "\nViewport Content", seo.ViewPort) // Change this line to refer to ViewPort
	}

	appendMetric(&builder, "\nRobots", boolToStr(seo.RobotsTxtPresent))

	// If you want to display alt texts, you can also add that here
	if len(seo.AltTexts) > 0 {
		builder.WriteString("\nAlt Texts:\n")
		for _, altText := range seo.AltTexts {
			builder.WriteString(" - " + altText + "\n")
		}
	}

	// If there are images without alt text, list them
	if len(seo.ImagesWithoutAltText) > 0 {
		builder.WriteString("\nImages without Alt Text:\n")
		for _, imgSrc := range seo.ImagesWithoutAltText {
			builder.WriteString(" - " + imgSrc + "\n")
		}
	}

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
