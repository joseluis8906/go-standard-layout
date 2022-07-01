package query

import (
	"context"
	"net/http"

	checkingaccount "github.com/joseluis8906/go-standard-layout/internal/domain/checking_account"
	"github.com/joseluis8906/go-standard-layout/internal/domain/customer"
)

type (
	GetCheckingAccounts struct {
		User string `json:"user_id"`
	}

	GetCheckingAccountsHandler struct {
		CheckingAccountFinder interface {
			GetAll(context.Context, customer.CustomerID) ([]checkingaccount.CheckingAccount, error)
		}
	}
)

func (g GetCheckingAccountsHandler) Do(ctx context.Context, query GetCheckingAccounts) ([]checkingaccount.CheckingAccount, error) {
	customerID, err := customer.ParseCurtomerID(query.User)
	if err != nil {
		return nil, err
	}

	return g.CheckingAccountFinder.GetAll(ctx, customerID)
}

func (g GetCheckingAccountsHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("unimplemented"))
}
