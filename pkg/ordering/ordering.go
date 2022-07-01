package ordering

const (
	Less    = Cmp(-1)
	Equal   = Cmp(0)
	Greater = Cmp(1)
)

type (
	Cmp int
)
