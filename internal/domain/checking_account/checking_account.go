package checkingaccount

import "github.com/joseluis8906/go-standard-layout/pkg/ddd/monetary"

type (
	CheckingAccount struct {
		number CheckingAccountNumber
		amount monetary.Amount
	}
)

func New() (CheckingAccount, error) {
	amount, err := monetary.NewAmount(0, monetary.USD)
	if err != nil {
		return NoopCheckingAccount(), err
	}

	ca := CheckingAccount{
		number: NewCheckingAccountNumber(),
		amount: amount,
	}

	return ca, nil
}

func NoopCheckingAccount() CheckingAccount {
	return CheckingAccount{
		number: NoopCheckingAccountNumber(),
		amount: monetary.NoopAmount(),
	}
}

func (c CheckingAccount) Number() CheckingAccountNumber {
	return c.number
}

func (c CheckingAccount) Amount() monetary.Amount {
	return c.amount
}
