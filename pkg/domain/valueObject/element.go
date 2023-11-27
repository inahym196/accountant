package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	ElementAssets      = "assets"
	ElementLiabilities = "liabilities"
	ElementEquaty      = "equaty"
	ElementExpense     = "expense"
	ElementIncome      = "income"
)

type element struct{ value string }

func (e element) String() string {
	return e.value
}

func NewElement(value string) (*element, error) {
	elements := []string{ElementAssets, ElementLiabilities, ElementEquaty, ElementExpense, ElementIncome}
	if !slices.Contains(elements, value) {
		return nil, fmt.Errorf("%s is not element", value)
	}
	return &element{value}, nil
}
