package database

import (
	"time"

	"github.com/zachlatta/angelhack/model"
)

func GetJournals(userID int64) ([]*model.Journal, error) {
	journals := []*model.Journal{}
	err := db.Select(&journals, "SELECT * FROM journals WHERE user_id=$1 ORDER BY id", userID)
	if err != nil {
		return nil, err
	}

	return journals, nil
}

func GetJournal(id int64) (*model.Journal, error) {
	journal := model.Journal{}
	err := db.Get(&journal, "SELECT * FROM journals WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &journal, nil
}

func SaveJournal(j *model.Journal) error {
	if j.ID == 0 {
		j.Created = time.Now()
	}
	j.Updated = time.Now()

	err := db.QueryRowx("INSERT INTO journals (user_id, created, updated, name) VALUES ($1, $2, $3, $4) RETURNING id", j.UserID, j.Created, j.Updated, j.Name).Scan(&j.ID)
	if err != nil {
		return err
	}

	return nil
}
