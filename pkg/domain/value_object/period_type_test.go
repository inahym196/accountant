package domain_test

import (
	"testing"

	domain "github.com/inahym196/accountant/pkg/domain/value_object"
)

func TestPeriodType(t *testing.T) {
	OKtests := []string{
		domain.PeriodDuration,
		domain.PeriodInstant,
	}
	for _, tt := range OKtests {
		_, err := domain.NewPeriodType(tt)
		if err != nil {
			t.Error(err)
		}
	}
	NGtests := []string{
		"test",
		"Duration",
	}
	for _, tt := range NGtests {
		_, err := domain.NewPeriodType(tt)
		if err == nil {
			t.Error("invalid status")
		}
	}
}
