package models

// PerformanceMetrics struct
type PerformanceMetrics struct {
	LCP                  float32 `json:"lcp"`                    // Largest Contentful Paint (LCP)
	FCP                  float32 `json:"fcp"`                    // First Contentful Paint (FCP)
	CLS                  float32 `json:"cls"`                    // Cumulative Layout Shift (CLS)
	INP                  float32 `json:"inp"`                    // Interaction to Next Paint (INP)
	TBT                  float32 `json:"tbt"`                    // Total Blocking Time (TBT)
	SpeedIndex           float32 `json:"speed-index"`            // Speed Index
	PageWeight           float32 `json:"page-weight"`            // Total page weight in KB or MB
	JavaScriptWeight     float32 `json:"javascript-weight"`      // JavaScript weight
	CSSWeight            float32 `json:"css-weight"`             // CSS weight
	HTTP2                bool    `json:"http2"`                  // HTTP/2 usage
	TextCompression      bool    `json:"text-compression"`       // GZIP/Brotli compression enabled
	CDNEnabled           bool    `json:"cdn-enabled"`            // Content Delivery Network enabled
	CacheControlHeaders  bool    `json:"cache-control-headers"`  // Cache-Control headers present
	KeepAliveConnections bool    `json:"keep-alive-connections"` // Keep-alive connections
	MinifiedJavaScript   bool    `json:"minified-javascript"`    // JavaScript minification
	MinifiedCSS          bool    `json:"minified-css"`           // CSS minification
}
