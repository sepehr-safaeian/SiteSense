package internal

import (
	"SiteSense/internal/collectors"
	"SiteSense/internal/process"
	"SiteSense/pkg/logger"
	"SiteSense/pkg/models"
	"fmt"
	"os"
	"strings"
)

// Collector triggers the collection process for a given URL
func Collector(url string) {
	seoData, collectedURL := collectors.SeoCollector(url)

	// Validate the collected data
	if err := validateSeoData(seoData, collectedURL); err != nil {
		log := logger.NewCustomLogger()
		log.Error(err.Error())
		return
	}

	// Create a file and save the collected data
	if err := createFile(collectedURL, seoData); err != nil {
		log := logger.NewCustomLogger()
		log.Error(err.Error())
	}
}

// validateSeoData checks if the SEO data contains a valid title
func validateSeoData(seoData models.SEOMetrics, collectedURL string) error {
	if !seoData.MetaTitlePresent {
		return fmt.Errorf("no H1 tag found for the URL: %s", collectedURL)
	}
	return nil
}

// createFile generates a new file based on the URL, storing the collected data in the ./data directory
func createFile(url string, seoData models.SEOMetrics) error {
	log := logger.NewCustomLogger()

	// Ensure the ./data directory exists
	if err := os.MkdirAll("./data", os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Convert the URL to a valid filename
	fileName := sanitizeFileName(url) + ".txt"
	filePath := "./data/" + fileName

	// Create or open the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Write collected SEO data to the file
	content := formatSeoData(seoData)
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	log.Info("File created successfully: " + filePath)
	return nil
}

// sanitizeFileName converts a URL to a safe and valid filename
func sanitizeFileName(url string) string {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	replacer := strings.NewReplacer(
		"/", "-",
		":", "-",
		"?", "-",
		"&", "-",
		"=", "-",
		"#", "-",
	)
	return replacer.Replace(url)
}

// formatSeoData formats the SEO data into a string for file output
func formatSeoData(seo models.SEOMetrics) string {
	return process.SeoProcess(seo)
}
