package domain_test

import (
	"testing"

	vo "github.com/inahym196/accountant/pkg/domain/value_object"
)

func NewAccountItem(title string, japanese_title string, period string, element string) vo.AccountItem {
	ai, _ := vo.NewAccountItem(title, japanese_title, period, element)
	return ai
}

func TestAccountItemEqual(t *testing.T) {
	ai := NewAccountItem("test", "test", vo.PeriodInstant, vo.ElementAssets)
	OKtests := []struct {
		name string
		ai   vo.AccountItem
		want bool
	}{
		{
			ai:   NewAccountItem("test", "test", vo.PeriodInstant, vo.ElementAssets),
			want: true,
		},
	}
	for _, tt := range OKtests {
		t.Run("match", func(t *testing.T) {
			if actual := ai.Equal(tt.ai); actual != tt.want {
				t.Errorf("invalid equal result: %v, want %v", actual, tt.want)
			}
		})
	}

	NGtests := []struct {
		name string
		ai   vo.AccountItem
		want bool
	}{
		{
			ai:   NewAccountItem("unmatch", "test", vo.PeriodInstant, vo.ElementAssets),
			want: false,
		},
		{
			ai:   NewAccountItem("test", "test", vo.PeriodDuration, vo.ElementAssets),
			want: false,
		},
		{
			ai:   NewAccountItem("test", "test", vo.PeriodDuration, vo.ElementEquaty),
			want: false,
		},
	}
	for _, tt := range NGtests {
		t.Run("unmatch", func(t *testing.T) {
			if actual := ai.Equal(tt.ai); actual != tt.want {
				t.Errorf("invalid equal result: %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestNewAccountItem(t *testing.T) {
	OKtests := []struct {
		title          string
		japanese_title string
		period_type    string
		element        string
		want           vo.AccountItem
	}{
		{
			title:          "test",
			japanese_title: "test",
			period_type:    "instant",
			element:        "assets",
			want:           NewAccountItem("test", "test", vo.PeriodInstant, vo.ElementAssets),
		},
	}
	for _, tt := range OKtests {
		t.Run("OK", func(t *testing.T) {
			ai, err := vo.NewAccountItem(tt.title, tt.japanese_title, tt.period_type, tt.element)
			if err != nil {
				t.Error(err)
			}
			if !ai.Equal(tt.want) {
				t.Errorf("invalid value: %v, want %v", ai, tt.want)
			}
		})
	}

	NGtests := []struct {
		title          string
		japanese_title string
		period_type    string
		element        string
		want           vo.AccountItem
	}{
		{
			title:          "test",
			japanese_title: "test",
			period_type:    "invalid-value",
			element:        "assets",
		},
		{
			title:          "test",
			japanese_title: "test",
			period_type:    "instant",
			element:        "invalid-value",
		},
	}
	for _, tt := range NGtests {
		t.Run("NG", func(t *testing.T) {
			ai, _ := vo.NewAccountItem(tt.title, tt.japanese_title, tt.period_type, tt.element)
			if ai != nil {
				t.Errorf("invalid status: %v, want nil", ai)
			}
		})
	}
}
