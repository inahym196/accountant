package domain_test

import (
	"testing"

	vo "github.com/inahym196/accountant/pkg/domain/value_object"
)

func TestAccountItemEqual(t *testing.T) {
	p1, _ := vo.NewPeriodType(vo.PeriodInstant)
	e1, _ := vo.NewElement(vo.ElementAssets)
	ai := vo.AccountItem{"test", "test", *p1, *e1}

	p2, _ := vo.NewPeriodType(vo.PeriodInstant)
	p3, _ := vo.NewPeriodType(vo.PeriodDuration)
	e2, _ := vo.NewElement(vo.ElementAssets)

	tests := []struct {
		name string
		ai   vo.AccountItem
		want bool
	}{
		{
			name: "match 1",
			ai:   ai,
			want: true,
		},
		{
			name: "match 2",
			ai:   vo.AccountItem{"test", "test", *p2, *e2},
			want: true,
		},
		{
			name: "unmatch 1",
			ai:   vo.AccountItem{"unmatch", "test", *p2, *e2},
			want: false,
		},
		{
			name: "unmatch 2",
			ai:   vo.AccountItem{"test", "test", *p3, *e2},
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
		if !ai.Equal(*tt.want) {
			t.Errorf("invalid value: %v, want %v", ai, tt.want)
		}
	}
}
