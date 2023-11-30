package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	SubjectBS     = "bs"
	SubjectPL     = "pl"
	BalanceDebit  = "debit"
	BalanceCredit = "credit"
)

type subject struct{ value string }

func NewSubject(value string) (*subject, error) {
	subjects := []string{SubjectBS, SubjectPL}
	if !slices.Contains(subjects, value) {
		return nil, fmt.Errorf("%s is not subject", value)
	}
	return &subject{value}, nil
}

func (s subject) String() string { return s.value }

type balance struct{ value string }

func NewBalance(value string) (*balance, error) {
	balances := []string{BalanceDebit, BalanceCredit}
	if !slices.Contains(balances, value) {
		return nil, fmt.Errorf("%s is not balance", value)
	}
	return &balance{value}, nil
}

func (b balance) String() string { return b.value }

type AccountItem struct {
	subject
	name   string
	jpname string
	periodType
	balance
}

func (item AccountItem) GetSubject() subject       { return item.subject }
func (item AccountItem) GetName() string           { return item.name }
func (item AccountItem) GetJPName() string         { return item.jpname }
func (item AccountItem) GetPeriodType() periodType { return item.periodType }
func (item AccountItem) GetBalance() balance       { return item.balance }

func NewAccountItem(subject string, name string, jpname string, period_type string, balance string) (*AccountItem, error) {
	subj, err := NewSubject(subject)
	if err != nil {
		return nil, err
	}

	period, err := NewPeriodType(period_type)
	if err != nil {
		return nil, err
	}

	b, err := NewBalance(balance)
	if err != nil {
		return nil, err
	}
	return &AccountItem{
		subject:    *subj,
		name:       name,
		jpname:     jpname,
		periodType: *period,
		balance:    *b,
	}, nil
}

type Account struct {
	AccountItem
	Amount uint32
}
