package database

import (
	"database/sql"
	"log"

	domain "github.com/inahym196/accountant/pkg/domain/value_object"
	"github.com/inahym196/accountant/pkg/usecase"
)

type accountItemDatabase struct {
	DB *sql.DB
}

func NewAccountItemRepository(db *sql.DB) usecase.AccountItemRepository {
	return &accountItemDatabase{DB: db}
}

func (repo accountItemDatabase) FindByTitle(title string) (*domain.AccountItem, error) {
	var jp_title, period, element string
	err := repo.DB.
		QueryRow("select japanese_title, period_type, element from account_item where title = ?", title).
		Scan(&jp_title, &period, &element)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}
	ai, err := domain.NewAccountItem(title, jp_title, period, element)
	if err != nil {
		return nil, err
	}
	return &ai, nil
}
func (repo accountItemDatabase) Save(ai domain.AccountItem) error {
	_, err := repo.DB.
		Exec(
			"replace into account_item (title,japanese_title,period_type,element) values(?,?,?,?)",
			ai.GetTitle(),
			ai.GetJapaneseTitle(),
			ai.GetPeriodType().String(),
			ai.GetElement().String(),
		)
	if err != nil {
		return err
	}
	return nil
}
