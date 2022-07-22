package fluentd

type (
	config struct {
		fluentdHost  string
		fluentdPort  int
		fluentdAsync bool
	}

	// OptionFunc ...
	OptionFunc func(*config)
)

// Host ...
func Host(opt string) OptionFunc {
	return func(c *config) {
		c.fluentdHost = opt
	}
}

// Port ...
func Port(opt int) OptionFunc {
	return func(c *config) {
		c.fluentdPort = opt
	}
}

// Async ...
func Async(opt bool) OptionFunc {
	return func(c *config) {
		c.fluentdAsync = opt
	}
}
