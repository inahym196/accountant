package domain_test

import (
	"testing"

	domain "github.com/inahym196/accountant/pkg/domain/value_object"
)

func TestElement(t *testing.T) {
	OKtests := []string{
		domain.ElementAssets,
		domain.ElementEquaty,
		domain.ElementExpense,
		domain.ElementIncome,
		domain.ElementLiabilities,
	}
	for _, tt := range OKtests {
		_, err := domain.NewElement(tt)
		if err != nil {
			t.Error(err)
		}
	}
	NGtests := []string{
		"test",
	}
	for _, tt := range NGtests {
		_, err := domain.NewElement(tt)
		if err == nil {
			t.Error("invalid status")
		}
	}
}
