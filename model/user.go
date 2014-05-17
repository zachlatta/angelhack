package model

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	ErrInvalidFirstName = errors.New("invalid first name")
	ErrInvalidLastName  = errors.New("invalid last name")
	ErrInvalidEmail     = errors.New("invalid email address")
	ErrInvalidPassword  = errors.New("invalid password")
)

var regexpEmail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

type User struct {
	ID        int64     `db:"id"         json:"id"`
	Created   time.Time `db:"created"    json:"created"`
	Updated   time.Time `db:"updated"    json:"updated"`
	FirstName string    `db:"first_name" json:"firstName"`
	LastName  string    `db:"last_name"  json:"lastName"`
	Email     string    `db:"email"      json:"email"`
	Password  string    `db:"password"   json:"-"`
}

type RequestUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUser(jsonReader io.Reader) (*User, error) {
	var rU RequestUser
	if err := json.NewDecoder(jsonReader).Decode(&rU); err != nil {
		return nil, err
	}

	if err := rU.validate(); err != nil {
		return nil, err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(rU.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		FirstName: rU.FirstName,
		LastName:  rU.LastName,
		Email:     rU.Email,
		Password:  string(b),
	}

	return &user, nil
}

func (u *RequestUser) validate() error {
	switch {
	case len(u.FirstName) == 0:
		return ErrInvalidFirstName
	case len(u.FirstName) >= 255:
		return ErrInvalidFirstName
	case len(u.LastName) == 0:
		return ErrInvalidLastName
	case len(u.LastName) >= 255:
		return ErrInvalidLastName
	case len(u.Email) >= 255:
		return ErrInvalidEmail
	case regexpEmail.MatchString(u.Email) == false:
		return ErrInvalidEmail
	case len(u.Password) < 6:
		return ErrInvalidPassword
	case len(u.Password) > 256:
		return ErrInvalidPassword
	default:
		return nil
	}
}

// ComparePassword compares the supplied password to the user password stored
// in the database.
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
