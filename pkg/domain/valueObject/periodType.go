package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	PeriodDuration = "duration"
	PeriodInstant  = "instant"
)

type periodType struct{ value string }

func (p periodType) String() string {
	return p.value
}

func NewPeriodType(value string) (*periodType, error) {
	periodTypes := []string{PeriodDuration, PeriodInstant}
	if !slices.Contains(periodTypes, value) {
		return nil, fmt.Errorf("%s is not periodType", value)
	}
	return &periodType{value}, nil
}
