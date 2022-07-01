package grpc

type (
	config struct {
		host string
		port int
	}

	OptionFunc func(*config)
)

func Host(opt string) OptionFunc {
	return func(cnf *config) {
		cnf.host = opt
	}
}

func Port(opt int) OptionFunc {
	return func(cnf *config) {
		cnf.port = opt
	}
}
