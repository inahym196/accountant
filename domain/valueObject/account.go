package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	PeriodDuration     = "duration"
	PeriodInstant      = "instant"
	BalanceDebit       = "debit"
	BalanceCredit      = "credit"
	ElementAssets      = "assets"
	ElementLiabilities = "liabilities"
	ElementEquaty      = "equaty"
	ElementExpense     = "expense"
	ElementIncome      = "income"
)

type PeriodType struct{ value string }

func NewPeriodType(periodType string) (*PeriodType, error) {
	periodTypes := []string{PeriodDuration, PeriodInstant}
	if slices.Contains(periodTypes, periodType) {
		return nil, fmt.Errorf("%s is not periodType", periodType)
	}
	return &PeriodType{periodType}, nil
}

type Balance struct{ value string }

func NewBalance(balance string) (*Balance, error) {
	balances := []string{BalanceCredit, BalanceDebit}
	if !slices.Contains(balances, balance) {
		return nil, fmt.Errorf("%s is not balance", balance)
	}
	return &Balance{balance}, nil
}

type Element struct{ value string }

func NewElement(element string) (*Element, error) {
	elements := []string{ElementAssets, ElementLiabilities, ElementEquaty, ElementExpense, ElementIncome}
	if !slices.Contains(elements, element) {
		return nil, fmt.Errorf("%s is not element", element)
	}
	return &Element{element}, nil
}

type AccountItem struct {
	Title         string
	JapaneseTitle string
	PeriodType    PeriodType
	Element       Element
}

type Amount uint32

type Account struct {
	AccountItem AccountItem
	Amount      Amount
}

type JournalID struct{ value uint32 }

func NewJournalID() *JournalID {
	return &JournalID{1}
}
