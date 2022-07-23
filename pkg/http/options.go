package http

type (
	config struct {
		allowedOrigins   []string
		allowedMethods   []string
		allowedHeaders   []string
		exposedHeaders   []string
		allowCredentials bool
		maxAge           int
	}

	// OptionFunc ...
	OptionFunc func(*config)
)

// AllowedOrigins ...
func AllowedOrigins(opt []string) OptionFunc {
	return func(c *config) {
		c.allowedOrigins = opt
	}
}

// AllowdMethods ...
func AllowdMethods(opt []string) OptionFunc {
	return func(c *config) {
		c.allowedMethods = opt
	}
}

// AllowedHeaders ...
func AllowedHeaders(opt []string) OptionFunc {
	return func(c *config) {
		c.allowedHeaders = opt
	}
}

// ExposedHeaders ...
func ExposedHeaders(opt []string) OptionFunc {
	return func(c *config) {
		c.exposedHeaders = opt
	}
}

// AllowCredentials ...
func AllowCredentials(opt bool) OptionFunc {
	return func(c *config) {
		c.allowCredentials = opt
	}
}

// MaxAge ...
func MaxAge(opt int) OptionFunc {
	return func(c *config) {
		c.maxAge = opt
	}
}
