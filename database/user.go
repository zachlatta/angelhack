package database

import (
	"database/sql"
	"time"

	"github.com/zachlatta/angelhack/model"
)

// GetUser gets the user from the database with the provided ID.
func GetUser(id int64) (*model.User, error) {
	user := model.User{}
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail gets the user from the database with the provided email.
func GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SaveUser saves the provided user to the database. If the user is a new
// user, then the user.Created field is set to the current time. The
// user.Updated field is set to the current time regardless. The user's
// information is also updated from the DB.
//
// If the user is a new user, then SaveUser also verifies that the user's
// email is unique and returns model.ErrInvalidEmail accordingly.
func SaveUser(u *model.User) error {
	if u.ID == 0 {
		// TODO: Should do this in the user.validate() method, but with the
		// current architecture of the application, that causes an import cycle.
		//
		// Check if email is unique.
		_, err := GetUserByEmail(u.Email)
		if err != sql.ErrNoRows {
			return model.ErrInvalidEmail
		}

		u.Created = time.Now()
	}
	u.Updated = time.Now()

	err := db.QueryRowx("INSERT INTO users (created, updated, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", u.Created, u.Updated, u.FirstName, u.LastName, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}
