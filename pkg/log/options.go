package log

type (
	config struct {
		fluentdHost  string
		fluentdPort  int
		fluentdAsync bool
		formatter    string
		level        int
		caller       bool
		env          string
	}

	OptionFunc func(*config)
)

func FluentdHost(opt string) OptionFunc {
	return func(c *config) {
		c.fluentdHost = opt
	}
}

func FluentdPort(opt int) OptionFunc {
	return func(c *config) {
		c.fluentdPort = opt
	}
}

func FluentdAsync(opt bool) OptionFunc {
	return func(c *config) {
		c.fluentdAsync = opt
	}
}

func Formatter(opt string) OptionFunc {
	return func(c *config) {
		c.formatter = opt
	}
}

func Level(opt int) OptionFunc {
	return func(c *config) {
		c.level = opt
	}
}

func Caller(opt bool) OptionFunc {
	return func(c *config) {
		c.caller = opt
	}
}

func Env(opt string) OptionFunc {
	return func(c *config) {
		c.env = opt
	}
}
