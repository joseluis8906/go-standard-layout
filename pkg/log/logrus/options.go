package logrus

type (
	config struct {
		formatter string
		level     int
		caller    bool
		env       string
	}

	// OptionFunc ...
	OptionFunc func(*config)
)

// Formatter ...
func Formatter(opt string) OptionFunc {
	return func(c *config) {
		c.formatter = opt
	}
}

// Level ...
func Level(opt int) OptionFunc {
	return func(c *config) {
		c.level = opt
	}
}

// Caller ...
func Caller(opt bool) OptionFunc {
	return func(c *config) {
		c.caller = opt
	}
}

// Env ...
func Env(opt string) OptionFunc {
	return func(c *config) {
		c.env = opt
	}
}
