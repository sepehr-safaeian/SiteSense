package models

// UsabilityMetrics struct
type UsabilityMetrics struct {
	ResponsiveDesign     bool `json:"responsive-design"`      // Responsive design for mobile devices
	TouchTargetsSize     bool `json:"touch-targets-size"`     // Large enough touch targets
	NavigationEase       bool `json:"navigation-ease"`        // Ease of navigation across the site
	BreadcrumbsEnabled   bool `json:"breadcrumbs-enabled"`    // Breadcrumb navigation enabled
	SearchFunctionality  bool `json:"search-functionality"`   // Search bar available and functional
	LoadingIndicators    bool `json:"loading-indicators"`     // Loading indicators for content-heavy areas
	ErrorPagesCustom     bool `json:"error-pages-custom"`     // Custom 404 and error pages
	FormFieldLabels      bool `json:"form-field-labels"`      // Form field labels are clear and descriptive
	MobileGestureSupport bool `json:"mobile-gesture-support"` // Mobile gesture support for swipe/zoom
	ScrollBarVisibility  bool `json:"scrollbar-visibility"`   // Scrollbars are visible when needed
	ConsistentStyling    bool `json:"consistent-styling"`     // Consistent design and styling across the site
}
