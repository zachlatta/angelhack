package handler

import (
	"errors"
	"net/http"

	"github.com/zachlatta/angelhack/database"
	"github.com/zachlatta/angelhack/model"
)

var (
	ErrUnauthorized = errors.New("user not authorized")
)

func Entries(w http.ResponseWriter, r *http.Request, u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	entries, err := database.GetEntries(u.ID)
	if err != nil {
		return &AppError{err, "error getting entries",
			http.StatusInternalServerError}
	}

	return renderJSON(w, entries, http.StatusOK)
}

func CreateEntry(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	defer r.Body.Close()
	entry, err := model.NewEntry(r.Body, u.ID)
	if err != nil {
		return &AppError{err, err.Error(), http.StatusBadRequest}
	}

	err = database.SaveEntry(entry)
	if err != nil {
		return &AppError{err, "error saving entry",
			http.StatusInternalServerError}
	}

	return renderJSON(w, entry, http.StatusOK)
}
