package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

type Journal struct {
	ID      int64     `db:"id" json:"id"`
	UserID  int64     `db:"user_id" json:"userID"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
	Name    string    `db:"name" json:"name"`
}

func NewJournal(jsonReader io.Reader, userID int64) (*Journal, error) {
	var journal Journal
	if err := json.NewDecoder(jsonReader).Decode(&journal); err != nil {
		return nil, err
	}

	journal.ID = 0

	journal.UserID = userID

	if err := journal.validate(); err != nil {
		return nil, err
	}

	return &journal, nil
}

func (j *Journal) validate() error {
	switch {
	case len(j.Name) < 1 || len(j.Name) > 50:
		return ErrInvalidName
	default:
		return nil
	}
}
