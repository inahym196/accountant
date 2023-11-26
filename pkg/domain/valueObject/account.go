package domain

type AccountItem interface {
	Equal(other AccountItem) bool
	GetTitle() string
	GetJapaneseTitle() string
	GetPeriodType() PeriodType
	GetElement() Element
}

type accountItem struct {
	title         string
	japaneseTitle string
	periodType    PeriodType
	element       Element
}

func (item accountItem) GetTitle() string          { return item.title }
func (item accountItem) GetJapaneseTitle() string  { return item.japaneseTitle }
func (item accountItem) GetPeriodType() PeriodType { return item.periodType }
func (item accountItem) GetElement() Element       { return item.element }

func (item accountItem) Equal(other AccountItem) bool {
	return item.title == other.GetTitle() &&
		item.japaneseTitle == other.GetJapaneseTitle() &&
		item.periodType == other.GetPeriodType() &&
		item.element == other.GetElement()
}

func NewAccountItem(title string, japanese_title string, period_type string, element string) (AccountItem, error) {
	period, err := NewPeriodType(period_type)
	if err != nil {
		return nil, err
	}
	elm, err := NewElement(element)
	if err != nil {
		return nil, err
	}
	return accountItem{
		title:         title,
		japaneseTitle: japanese_title,
		periodType:    period,
		element:       elm,
	}, nil
}

type Account struct {
	AccountItem
	Amount uint32
}
