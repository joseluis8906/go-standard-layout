package checkingaccount

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/joseluis8906/go-standard-layout/pkg/errors"
)

const (
	ErrInvalidCheckingAccountNumber         = errors.Error("invalid checking account number")
	ErrInvalidCheckingAccountNumberBankCode = errors.Error("invalid checking account nummber bank code")
	ErrInvalidCheckingAccountNumberType     = errors.Error("invalid checking account nummber type")
	ErrInvalidCheckingAccountNumberRange    = errors.Error("invalid checking account nummber range")
)

const (
	nilCheckingAccountVal = "<nil>"
	bankCode              = 3729
	checkingAccountType   = 28325
)

type (
	CheckingAccountNumber struct {
		val string
	}
)

func NewCheckingAccountNumber() CheckingAccountNumber {
	accountNumber := rand.Int31n((9999 - 1000) + 1000)

	return CheckingAccountNumber{val: fmt.Sprintf("%d-%d-%d", bankCode, checkingAccountType, accountNumber)}
}

func NoopCheckingAccountNumber() CheckingAccountNumber {
	return CheckingAccountNumber{val: nilCheckingAccountVal}
}

func ParseCheckingAccountNumber(val string) (CheckingAccountNumber, error) {
	raw := strings.Split(val, "-")
	if len(raw) != 3 {
		return NoopCheckingAccountNumber(), ErrInvalidCheckingAccountNumber
	}

	localBankCode, err := strconv.Atoi(raw[0])
	if err != nil {
		return NoopCheckingAccountNumber(), err
	}

	localChekingAccountType, err := strconv.Atoi(raw[1])
	if err != nil {
		return NoopCheckingAccountNumber(), err
	}

	checkingAccountNumber, err := strconv.Atoi(raw[2])
	if err != nil {
		return NoopCheckingAccountNumber(), err
	}

	if localBankCode != bankCode {
		return NoopCheckingAccountNumber(), ErrInvalidCheckingAccountNumberBankCode
	}

	if localChekingAccountType != checkingAccountType {
		return NoopCheckingAccountNumber(), ErrInvalidCheckingAccountNumberType
	}

	if checkingAccountNumber > 9999 || checkingAccountNumber < 1000 {
		return NoopCheckingAccountNumber(), ErrInvalidCheckingAccountNumberRange
	}

	return CheckingAccountNumber{val}, nil
}

func (can CheckingAccountNumber) IsZero() bool {
	return can.val == nilCheckingAccountVal
}

func (can CheckingAccountNumber) String() string {
	return can.val
}
