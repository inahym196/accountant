package domain

import (
	"time"

	valueObject "github.com/inahym196/accountant/pkg/domain/value_object"
)

type JournalEntry struct {
	Id          valueObject.JournalID
	Date        time.Time
	DebitEntry  []valueObject.Account
	CreditEntry []valueObject.Account
}
