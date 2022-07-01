package redis

type (
	config struct {
		host   string
		port   int
		passwd string
		db     int
	}

	OptionFunc func(*config)
)

func Host(opt string) OptionFunc {
	return func(c *config) {
		c.host = opt
	}
}

func Port(opt int) OptionFunc {
	return func(c *config) {
		c.port = opt
	}
}

func Passwd(opt string) OptionFunc {
	return func(c *config) {
		c.passwd = opt
	}
}

func DB(opt int) OptionFunc {
	return func(c *config) {
		c.db = opt
	}
}
