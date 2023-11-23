package usecase

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
	repo IAccountItemRepository
}

func NewAccountItemInteractor(repo IAccountItemRepository) IAccountItemUseCase {
	return &AccountItemInteractor{repo}
}

type IAccountItemRepository interface {
	FindByTitle(title string) (*AccountItemDTO, error)
	Save(ai AccountItemDTO) error
}

func (i AccountItemInteractor) FindByTitle(title string) (*AccountItemDTO, error) {
	dto, err := i.repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
func (i AccountItemInteractor) Save(dto AccountItemDTO) error {
	return i.repo.Save(dto)
}
