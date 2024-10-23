package internal

import (
	"SiteSense/internal/collectors"
	"SiteSense/internal/process"
	"SiteSense/pkg/logger"
	"SiteSense/pkg/models"
	"os"
	"strings"
)

// Collector triggers the collection process for a given URL
func Collector(url string) {
	// Collect SEO data from the URL
	seoData := collectors.SeoCollector(url)

	// Create a file and save the collected data
	createFile(url, seoData)
}

// createFile generates a new file based on the URL, storing the collected data in the ./data directory
func createFile(url string, seoData models.SEOMetrics) {
	log := logger.NewCustomLogger()

	// Ensure the ./data directory exists
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Error("Error creating directory: " + err.Error())
		return
	}

	// Convert the URL to a valid filename
	fileName := sanitizeFileName(url) + ".txt"
	filePath := "./data/" + fileName

	// Create or open the file
	file, err := os.Create(filePath)
	if err != nil {
		log.Error("Error creating file: " + err.Error())
		return
	}
	defer file.Close()

	// Write collected SEO data to the file
	content := formatSeoData(seoData)
	if _, err := file.WriteString(content); err != nil {
		log.Error("Error writing to file: " + err.Error())
		return
	}

	log.Info("File created successfully: " + filePath)
}

// sanitizeFileName converts a URL to a safe and valid filename
func sanitizeFileName(url string) string {
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	return strings.NewReplacer(
		"/", "-",
		":", "-",
		"?", "-",
		"&", "-",
		"=", "-",
		"#", "-",
	).Replace(url)
}

// formatSeoData formats the SEO data into a string for file output
func formatSeoData(seo models.SEOMetrics) string {
	return process.SeoProcess(seo)
}
