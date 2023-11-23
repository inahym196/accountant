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
	FindByTitle(title string) (*AccountItemDTO, error)
	Save(ai AccountItemDTO) error
}

func (i accountItemInteractor) FindByTitle(title string) (*AccountItemDTO, error) {
	dto, err := i.repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
func (i accountItemInteractor) Save(dto AccountItemDTO) error {
	ai, err := domain.NewAccountItem(dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element)
	if err != nil {
		return err
	}
	repo_dto := AccountItemDTO{
		Title:         ai.GetTitle(),
		JapaneseTitle: ai.GetJapaneseTitle(),
		PeriodType:    ai.GetPeriodType().String(),
		Element:       ai.GetElement().String(),
	}
	return i.repo.Save(repo_dto)
}
