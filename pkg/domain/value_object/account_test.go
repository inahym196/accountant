package domain_test

import (
	"testing"

	vo "github.com/inahym196/accountant/pkg/domain/value_object"
)

func TestAccountItem(t *testing.T) {

	p, _ := vo.NewPeriodType(vo.PeriodInstant)
	e, _ := vo.NewElement(vo.ElementAssets)
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
			want:           &vo.AccountItem{"test", "test", *p, *e},
		},
	}
	for _, tt := range tests {
		ai, err := vo.NewAccountItem(tt.title, tt.japanese_title, tt.period_type, tt.element)
		if err != nil {
			t.Error(err)
		}
		if ai != tt.want {
			t.Errorf("invalid value: %v, want %v", ai, tt.want)
		}
	}
}
