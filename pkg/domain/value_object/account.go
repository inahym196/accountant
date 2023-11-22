package domain

type AccountItem interface {
	Equal(other AccountItem) bool
	GetTitle() string
	GetJapaneseTitle() string
	GetPeriodType() PeriodType
	GetElement() Element
}

type accountItem struct {
	AccountItem
	title         string
	japaneseTitle string
	periodType    PeriodType
	element       Element
}

func (ai accountItem) GetTitle() string          { return ai.title }
func (ai accountItem) GetJapaneseTitle() string  { return ai.japaneseTitle }
func (ai accountItem) GetPeriodType() PeriodType { return ai.periodType }
func (ai accountItem) GetElement() Element       { return ai.element }

func (ai accountItem) Equal(other AccountItem) bool {
	return ai.title == other.GetTitle() &&
		ai.japaneseTitle == other.GetJapaneseTitle() &&
		ai.periodType == other.GetPeriodType() &&
		ai.element == other.GetElement()
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
		periodType:    *period,
		element:       *elm,
	}, nil
}

type IAccountItemRepository interface {
	FindByTitle(title string) (AccountItem, error)
	Save(ai AccountItem) error
}

type Amount uint32

type Account struct {
	AccountItem AccountItem
	Amount      Amount
}
