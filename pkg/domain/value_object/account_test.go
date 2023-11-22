package domain_test

import (
	"testing"

	vo "github.com/inahym196/accountant/pkg/domain/value_object"
)

func NewPeriodType(period string) vo.PeriodType {
	p, _ := vo.NewPeriodType(period)
	return *p
}
func NewElement(element string) vo.Element {
	e, _ := vo.NewElement(element)
	return *e
}

func TestAccountItemEqual(t *testing.T) {
	ai := vo.AccountItem{"test", "test", NewPeriodType(vo.PeriodInstant), NewElement(vo.ElementAssets)}
	tests := []struct {
		name string
		ai   vo.AccountItem
		want bool
	}{
		{
			name: "match 1",
			ai:   vo.AccountItem{"test", "test", NewPeriodType(vo.PeriodInstant), NewElement(vo.ElementAssets)},
			want: true,
		},
		{
			name: "unmatch 1",
			ai:   vo.AccountItem{"unmatch", "test", NewPeriodType(vo.PeriodInstant), NewElement(vo.ElementAssets)},
			want: false,
		},
		{
			name: "unmatch 2",
			ai:   vo.AccountItem{"test", "test", NewPeriodType(vo.PeriodDuration), NewElement(vo.ElementAssets)},
			want: false,
		},
		{
			name: "unmatch 3",
			ai:   vo.AccountItem{"test", "test", NewPeriodType(vo.PeriodDuration), NewElement(vo.ElementEquaty)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := ai.Equal(tt.ai); actual != tt.want {
				t.Errorf("invalid equal result: %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestNewAccountItem(t *testing.T) {
	tests := []struct {
		title          string
		japanese_title string
		period_type    string
		element        string
		want           *vo.AccountItem
	}{
		{
			title:          "test",
			japanese_title: "test",
			period_type:    "instant",
			element:        "assets",
			want:           &vo.AccountItem{"test", "test", NewPeriodType(vo.PeriodInstant), NewElement(vo.ElementAssets)},
		},
	}
	for _, tt := range tests {
		ai, err := vo.NewAccountItem(tt.title, tt.japanese_title, tt.period_type, tt.element)
		if err != nil {
			t.Error(err)
		}
		if !ai.Equal(*tt.want) {
			t.Errorf("invalid value: %v, want %v", ai, tt.want)
		}
	}
}
