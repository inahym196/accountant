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

func (pt PeriodType) String() string {
	return pt.value
}

func NewPeriodType(periodType string) (*PeriodType, error) {
	periodTypes := []string{PeriodDuration, PeriodInstant}
	if !slices.Contains(periodTypes, periodType) {
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

func (elm Element) String() string {
	return elm.value
}

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

func NewAccountItem(title string, japanese_title string, period_type string, element string) (*AccountItem, error) {
	pt, err := NewPeriodType(period_type)
	if err != nil {
		return nil, err
	}
	elm, err := NewElement(element)
	if err != nil {
		return nil, err
	}
	return &AccountItem{
		Title:         title,
		JapaneseTitle: japanese_title,
		PeriodType:    *pt,
		Element:       *elm,
	}, nil
}

type IAccountItemRepository interface {
	FindByTitle(title string) (*AccountItem, error)
	Save(ai *AccountItem) error
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
