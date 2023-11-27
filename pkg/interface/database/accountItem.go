package database

import (
	"database/sql"
	"log"

	domain "github.com/inahym196/accountant/pkg/domain/valueObject"
	"github.com/inahym196/accountant/pkg/usecase"
)

type accountItemDatabase struct {
	DB *sql.DB
}

func NewAccountItemRepository(db *sql.DB) usecase.AccountItemRepository {
	return &accountItemDatabase{DB: db}
}

func (repo accountItemDatabase) FindByTitle(subject string, name string) (*domain.AccountItem, error) {
	var jpname, period, balance string
	err := repo.DB.
		QueryRow("select jpname, periodtype, balance from accountItem where subject = ? and name = ? ", subject, name).
		Scan(&jpname, &period, &balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}
	ai, err := domain.NewAccountItem(subject, name, jpname, period, balance)
	if err != nil {
		return nil, err
	}
	return ai, nil
}

func (repo accountItemDatabase) GetAll() ([]*domain.AccountItem, error) {
	rows, err := repo.DB.Query("select subject, name, jpname, periodtype, balance from accountItem")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var ais []*domain.AccountItem
	for rows.Next() {
		var subject, name, jpname, period, balance string
		if err := rows.Scan(&subject, &name, &jpname, &period, &balance); err != nil {
			return nil, err
		}
		ai, err := domain.NewAccountItem(subject, name, jpname, period, balance)
		if err != nil {
			return nil, err
		}
		ais = append(ais, ai)
	}
	return ais, err
}

func (repo accountItemDatabase) Save(ai *domain.AccountItem) error {
	_, err := repo.DB.
		Exec(
			"replace into accountItem (subject, name, jpname, periodtype, balance) values(?,?,?,?,?)",
			ai.GetSubject().String(),
			ai.GetName(),
			ai.GetJPName(),
			ai.GetPeriodType().String(),
			ai.GetBalance().String(),
		)
	if err != nil {
		return err
	}
	return nil
}
