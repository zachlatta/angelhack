package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var (
	ErrInvalidRating  = errors.New("invalid rating")
	ErrInvalidMessage = errors.New("invalid message")
)

type Entry struct {
	ID        int64     `db:"id"      json:"id"`
	JournalID int64     `db:"journal_id" json:"journalID"`
	Created   time.Time `db:"created" json:"created"`
	Updated   time.Time `db:"updated" json:"updated"`
	Rating    int       `db:"rating"  json:"rating"`
	Message   string    `db:"message" json:"message"`
}

func NewEntry(jsonReader io.Reader, journalID int64) (*Entry, error) {
	var entry Entry
	if err := json.NewDecoder(jsonReader).Decode(&entry); err != nil {
		return nil, err
	}

	entry.ID = 0

	entry.JournalID = journalID

	if err := entry.validate(); err != nil {
		return nil, err
	}

	return &entry, nil
}

func (u *Entry) validate() error {
	switch {
	case u.Rating < 1 || u.Rating > 5:
		return ErrInvalidRating
	case (len(u.Message) < 0 || len(u.Message) > 1024):
		return ErrInvalidMessage
	default:
		return nil
	}
}
