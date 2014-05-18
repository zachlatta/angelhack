package database

import (
	"time"

	"github.com/zachlatta/angelhack/model"
)

func GetJournalEntries(journalID int64) ([]model.Entry, error) {
	entries := []model.Entry{}
	err := db.Select(&entries, "SELECT * FROM entries WHERE journal_id=$1 ORDER BY id", journalID)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func GetEntry(id int64) (*model.Entry, error) {
	user := model.Entry{}
	err := db.Get(&user, "SELECT * FROM entries WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteEntry(id int64) error {
	_, err := db.Exec("DELETE FROM entries WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func SaveEntry(e *model.Entry) error {
	if e.ID == 0 {
		e.Created = time.Now()
	}
	e.Updated = time.Now()

	err := db.QueryRowx("INSERT INTO entries (journal_id, created, updated, rating, message) VALUES ($1, $2, $3, $4, $5) RETURNING id", e.JournalID, e.Created, e.Updated, e.Rating, e.Message).Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}
