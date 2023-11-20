package domain

import (
	"time"

	vo "github.com/inahym196/accountant/pkg/domain/value_object"
)

type JournalEntry struct {
	Id          int
	Date        time.Time
	DebitEntry  []vo.Account
	CreditEntry []vo.Account
}
