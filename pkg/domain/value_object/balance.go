package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	BalanceDebit  = "debit"
	BalanceCredit = "credit"
)

type Balance struct{ value string }

func NewBalance(value string) (*Balance, error) {
	balances := []string{BalanceCredit, BalanceDebit}
	if !slices.Contains(balances, value) {
		return nil, fmt.Errorf("%s is not balance", value)
	}
	return &Balance{value}, nil
}
