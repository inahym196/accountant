package domain

type JournalID struct{ value uint32 }

func NewJournalID() *JournalID {
	return &JournalID{1}
}
