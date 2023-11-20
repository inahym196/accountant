package sqlite

import (
	"database/sql"
	"log"

	domain "github.com/inahym196/accountant/pkg/domain/value_object"
)

type AccountItemRepository struct {
	DB *sql.DB
}

func NewAccountItemRepository(db *sql.DB) domain.IAccountItemRepository {
	return &AccountItemRepository{DB: db}
}

func (repo AccountItemRepository) FindByTitle(t string) (*domain.AccountItem, error) {
	var jp_title, period, element string
	err := repo.DB.QueryRow("select japanese_title, period_type, element from account_item where title = ?", t).Scan(&jp_title, &period, &element)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}
	ai, err := domain.NewAccountItem(t, jp_title, period, element)
	if err != nil {
		return nil, err
	}
	return ai, nil
}
func (repo AccountItemRepository) Save(ai *domain.AccountItem) error {
	_, err := repo.DB.Exec("replace into account_item (title,japanese_title,period_type,element) values(?,?,?,?)", ai.Title, ai.JapaneseTitle, ai.PeriodType, ai.Element)
	if err != nil {
		return err
	}
	return nil
}
