package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	PeriodDuration = "duration"
	PeriodInstant  = "instant"
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
