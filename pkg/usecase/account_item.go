package usecase

import domain "github.com/inahym196/accountant/pkg/domain/value_object"

type AccountItemDTO struct {
	Title         string
	JapaneseTitle string
	PeriodType    string
	Element       string
}

type AccountItemUseCase interface {
	FindByTitle(title string) (*AccountItemDTO, error)
	Save(AccountItemDTO) error
}

type accountItemInteractor struct {
	repo AccountItemRepository
}

func NewAccountItemInteractor(repo AccountItemRepository) AccountItemUseCase {
	return &accountItemInteractor{repo}
}

type AccountItemRepository interface {
	FindByTitle(title string) (*domain.AccountItem, error)
	Save(ai domain.AccountItem) error
}

func (i accountItemInteractor) FindByTitle(title string) (*AccountItemDTO, error) {
	ai, err := i.repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return &AccountItemDTO{
		Title:         (*ai).GetTitle(),
		JapaneseTitle: (*ai).GetJapaneseTitle(),
		PeriodType:    (*ai).GetPeriodType().String(),
		Element:       (*ai).GetElement().String(),
	}, nil
}

func (i accountItemInteractor) Save(dto AccountItemDTO) error {
	ai, err := domain.NewAccountItem(dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element)
	if err != nil {
		return err
	}
	return i.repo.Save(ai)
}
