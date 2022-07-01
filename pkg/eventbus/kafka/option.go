package kafka

type (
	config struct {
		bootstrapServers string
		groupID          string
		autoOffsetReset  string
		consumerTopics   []string
	}

	OptionFunc func(*config)
)

func BootstrapServers(opt string) OptionFunc {
	return func(c *config) {
		c.bootstrapServers = opt
	}
}

func GroupId(opt string) OptionFunc {
	return func(c *config) {
		c.groupID = opt
	}
}

func AutoOffsetReset(opt string) OptionFunc {
	return func(c *config) {
		c.autoOffsetReset = opt
	}
}

func ConsumerTopics(opt []string) OptionFunc {
	return func(c *config) {
		c.consumerTopics = opt
	}
}
