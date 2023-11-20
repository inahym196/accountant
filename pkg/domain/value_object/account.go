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

func (p PeriodType) String() string {
	return p.value
}

func NewPeriodType(value string) (*PeriodType, error) {
	periodTypes := []string{PeriodDuration, PeriodInstant}
	if !slices.Contains(periodTypes, value) {
		return nil, fmt.Errorf("%s is not periodType", value)
	}
	return &PeriodType{value}, nil
}

type Balance struct{ value string }

func NewBalance(value string) (*Balance, error) {
	balances := []string{BalanceCredit, BalanceDebit}
	if !slices.Contains(balances, value) {
		return nil, fmt.Errorf("%s is not balance", value)
	}
	return &Balance{value}, nil
}

type Element struct{ value string }

func (e Element) String() string {
	return e.value
}

func NewElement(value string) (*Element, error) {
	elements := []string{ElementAssets, ElementLiabilities, ElementEquaty, ElementExpense, ElementIncome}
	if !slices.Contains(elements, value) {
		return nil, fmt.Errorf("%s is not element", value)
	}
	return &Element{value}, nil
}

type AccountItem struct {
	Title         string
	JapaneseTitle string
	PeriodType    PeriodType
	Element       Element
}

func NewAccountItem(title string, japanese_title string, period_type string, element string) (*AccountItem, error) {
	period, err := NewPeriodType(period_type)
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
		PeriodType:    *period,
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
