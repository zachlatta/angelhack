package database

import (
	"time"

	"github.com/zachlatta/angelhack/model"
)

func GetEntries(userID int64) ([]model.Entry, error) {
	users := []model.Entry{}
	err := db.Select(&users, "SELECT * FROM entries WHERE user_id=$1 ORDER BY id", userID)
	if err != nil {
		return nil, err
	}

	return users, nil
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

	err := db.QueryRowx("INSERT INTO entries (user_id, created, updated, rating, message) VALUES ($1, $2, $3, $4, $5) RETURNING id", e.UserID, e.Created, e.Updated, e.Rating, e.Message).Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}
