package domain

import (
	"time"

	valueObject "github.com/inahym196/accountant/domain/valueObject"
)

type JournalEntry struct {
	Id          valueObject.JournalID
	Date        time.Time
	DebitEntry  []valueObject.Account
	CreditEntry []valueObject.Account
}
