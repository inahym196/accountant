package database

import (
	"database/sql"
	"log"

	"github.com/inahym196/accountant/pkg/usecase"
)

type AccountItemDatabase struct {
	DB *sql.DB
}

func NewAccountItemRepository(db *sql.DB) usecase.AccountItemRepository {
	return &AccountItemDatabase{DB: db}
}

func (repo AccountItemDatabase) FindByTitle(title string) (*usecase.AccountItemDTO, error) {
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
	return &usecase.AccountItemDTO{
		Title:         title,
		JapaneseTitle: jp_title,
		PeriodType:    period,
		Element:       element,
	}, nil
}
func (repo AccountItemDatabase) Save(dto usecase.AccountItemDTO) error {
	_, err := repo.DB.
		Exec(
			"replace into account_item (title,japanese_title,period_type,element) values(?,?,?,?)",
			dto.Title,
			dto.JapaneseTitle,
			dto.PeriodType,
			dto.Element,
		)
	if err != nil {
		return err
	}
	return nil
}
