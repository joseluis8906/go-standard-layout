package mongo

type (
	config struct {
		host   string
		port   int
		user   string
		passwd string
		db     string
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

func User(opt string) OptionFunc {
	return func(c *config) {
		c.user = opt
	}
}

func Passwd(opt string) OptionFunc {
	return func(c *config) {
		c.passwd = opt
	}
}

func DB(opt string) OptionFunc {
	return func(c *config) {
		c.db = opt
	}
}
