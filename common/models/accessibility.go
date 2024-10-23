package models

// AccessibilityMetrics struct
type AccessibilityMetrics struct {
	AltTextForImages     bool `json:"alt-text-for-images"`    // Alt text for images
	AriaAttributes       bool `json:"aria-attributes"`        // Proper ARIA attributes
	KeyboardNavigable    bool `json:"keyboard-navigable"`     // Full keyboard navigability
	TabOrderCorrect      bool `json:"tab-order-correct"`      // Logical tab order for keyboard navigation
	ColorContrast        bool `json:"color-contrast"`         // Sufficient color contrast between text and background
	FocusVisible         bool `json:"focus-visible"`          // Focus indicators for interactive elements
	LanguageDeclared     bool `json:"language-declared"`      // HTML lang attribute correctly set
	AccessibleForms      bool `json:"accessible-forms"`       // Forms are accessible with proper labels
	SkipToContentLink    bool `json:"skip-to-content-link"`   // "Skip to content" link available
	TextResizability     bool `json:"text-resizability"`      // Text can be resized without loss of content or functionality
	MediaWithCaptions    bool `json:"media-with-captions"`    // Videos and audio have captions or transcripts
	NoFlashingContent    bool `json:"no-flashing-content"`    // Avoids flashing content that may cause seizures
	HeadingsOrganized    bool `json:"headings-organized"`     // Correct use of heading levels (H1, H2, etc.)
	AccessibleSVG        bool `json:"accessible-svg"`         // Accessible SVGs with title and description
	FormErrorIndicators  bool `json:"form-error-indicators"`  // Clear error indicators in forms
	ScreenReaderFriendly bool `json:"screen-reader-friendly"` // Screen-reader compatibility
}
