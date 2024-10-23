package models

// SecurityMetrics struct
type SecurityMetrics struct {
	TLS_SSL                bool `json:"tls-ssl"`                  // TLS/SSL enabled
	HSTS                   bool `json:"hsts"`                     // HTTP Strict Transport Security enabled
	ContentSecurityPolicy  bool `json:"content-security-policy"`  // CSP headers in place
	XFrameOptions          bool `json:"x-frame-options"`          // X-Frame-Options header
	XContentTypeOptions    bool `json:"x-content-type-options"`   // X-Content-Type-Options header
	SubresourceIntegrity   bool `json:"subresource-integrity"`    // Subresource integrity for external scripts/styles
	CSPNonceUsed           bool `json:"csp-nonce-used"`           // CSP nonce for inline scripts/styles
	CookiesSecure          bool `json:"cookies-secure"`           // Secure flag on cookies
	CookiesHttpOnly        bool `json:"cookies-http-only"`        // HttpOnly flag on cookies
	AuthenticationSecure   bool `json:"authentication-secure"`    // Secure authentication mechanisms
	PasswordStorage        bool `json:"password-storage"`         // Strong password storage (e.g., bcrypt, scrypt)
	RateLimiting           bool `json:"rate-limiting"`            // Rate limiting implemented
	CORS                   bool `json:"cors"`                     // Proper Cross-Origin Resource Sharing policy
	AntiCSRFProtection     bool `json:"anti-csrf-protection"`     // Anti-CSRF tokens in forms and APIs
	SQLInjectionProtection bool `json:"sql-injection-protection"` // Protection against SQL injection
	XSSProtection          bool `json:"xss-protection"`           // Cross-site scripting protection
	DirectoryListing       bool `json:"directory-listing"`        // Directory listing disabled
	FirewallProtection     bool `json:"firewall-protection"`      // Web Application Firewall (WAF) in place
}
