package usecase

import (
	domain "github.com/inahym196/accountant/pkg/domain/value_object"
)

type AccountItemDTO struct {
	Title         string
	JapaneseTitle string
	PeriodType    string
	Element       string
}

type IAccountItemUseCase interface {
	FindByTitle(title string) (*AccountItemDTO, error)
	Save(AccountItemDTO) error
}

type AccountItemInteractor struct {
	repo domain.IAccountItemRepository
}

func NewAccountItemInteractor(repo domain.IAccountItemRepository) IAccountItemUseCase {
	return &AccountItemInteractor{repo}
}

func (i AccountItemInteractor) FindByTitle(title string) (*AccountItemDTO, error) {
	ai, err := i.repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return &AccountItemDTO{
		Title:         ai.GetTitle(),
		JapaneseTitle: ai.GetJapaneseTitle(),
		PeriodType:    ai.GetPeriodType().String(),
		Element:       ai.GetElement().String(),
	}, nil

}
func (i AccountItemInteractor) Save(dto AccountItemDTO) error {
	ai, err := domain.NewAccountItem(dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element)
	if err != nil {
		return err
	}
	return i.repo.Save(ai)
}
