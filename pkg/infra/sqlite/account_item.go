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

func (ai_repo AccountItemRepository) FindByTitle(title string) (*domain.AccountItem, error) {
	var t string
	var jt string
	var pt string
	var e string
	err := ai_repo.DB.QueryRow("select title,japanese_title, period_type, element from account_item where title = ?", title).
		Scan(&t, &jt, &pt, &e)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}
	ai, err := domain.NewAccountItem(t, jt, pt, e)
	if err != nil {
		return nil, err
	}
	return ai, nil
}
func (ai_repo AccountItemRepository) Save(ai *domain.AccountItem) error {
	_, err := ai_repo.DB.Exec("replace into account_item (title,japanese_title,period_type,element) values(?,?,?,?)", ai.Title, ai.JapaneseTitle, ai.PeriodType, ai.Element)
	if err != nil {
		return err
	}
	return nil
}
