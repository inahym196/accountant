package domain

type AccountItem struct {
	Title         string
	JapaneseTitle string
	PeriodType    PeriodType
	Element       Element
}

func (ai AccountItem) Equal(other AccountItem) bool {
	return ai.Title == other.Title &&
		ai.JapaneseTitle == other.JapaneseTitle &&
		ai.PeriodType == other.PeriodType &&
		ai.Element == other.Element
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
