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
	ai_repo domain.IAccountItemRepository
}

func NewAccountItemInteractor(repo domain.IAccountItemRepository) IAccountItemUseCase {
	return &AccountItemInteractor{ai_repo: repo}
}

func (ai_i AccountItemInteractor) FindByTitle(title string) (*AccountItemDTO, error) {
	ai, err := ai_i.ai_repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return &AccountItemDTO{
		Title:         ai.Title,
		JapaneseTitle: ai.JapaneseTitle,
		PeriodType:    ai.PeriodType.String(),
		Element:       ai.Element.String(),
	}, nil

}
func (ai_i AccountItemInteractor) Save(ai_dto AccountItemDTO) error {
	ai, err := domain.NewAccountItem(ai_dto.Title, ai_dto.JapaneseTitle, ai_dto.PeriodType, ai_dto.Element)
	if err != nil {
		return err
	}
	return ai_i.ai_repo.Save(ai)
}
