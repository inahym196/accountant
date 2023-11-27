package usecase

import domain "github.com/inahym196/accountant/pkg/domain/valueObject"

type AccountItemDTO struct {
	Subject    string
	Name       string
	JPName     string
	PeriodType string
	Balance    string
}

type AccountItemUseCase interface {
	FindByTitle(subject string, name string) (*AccountItemDTO, error)
	GetAll() ([]AccountItemDTO, error)
	Save(AccountItemDTO) error
}

type accountItemInteractor struct {
	repo AccountItemRepository
}

func NewAccountItemInteractor(repo AccountItemRepository) AccountItemUseCase {
	return accountItemInteractor{repo}
}

type AccountItemRepository interface {
	FindByTitle(subject string, name string) (*domain.AccountItem, error)
	GetAll() ([]*domain.AccountItem, error)
	Save(ai *domain.AccountItem) error
}

func (i accountItemInteractor) FindByTitle(subject string, name string) (*AccountItemDTO, error) {
	ai, err := i.repo.FindByTitle(subject, name)
	if err != nil {
		return nil, err
	}
	return &AccountItemDTO{
		Subject:    ai.GetSubject().String(),
		Name:       ai.GetName(),
		JPName:     ai.GetJPName(),
		PeriodType: ai.GetPeriodType().String(),
		Balance:    ai.GetBalance().String(),
	}, nil
}

func (i accountItemInteractor) GetAll() ([]AccountItemDTO, error) {
	ais, err := i.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var dtos []AccountItemDTO
	for _, ai := range ais {
		dto := AccountItemDTO{
			Subject:    ai.GetSubject().String(),
			Name:       ai.GetName(),
			JPName:     ai.GetJPName(),
			PeriodType: ai.GetPeriodType().String(),
			Balance:    ai.GetBalance().String(),
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

func (i accountItemInteractor) Save(dto AccountItemDTO) error {
	ai, err := domain.NewAccountItem(dto.Subject, dto.Name, dto.JPName, dto.PeriodType, dto.Balance)
	if err != nil {
		return err
	}
	return i.repo.Save(ai)
}
