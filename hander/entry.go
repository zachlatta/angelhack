package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zachlatta/angelhack/database"
	"github.com/zachlatta/angelhack/model"
)

var (
	ErrUnauthorized = errors.New("user not authorized")
)

func JournalEntries(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	vars := mux.Vars(r)
	journalID, err := strconv.ParseInt(vars["journalID"], 10, 64)
	if err != nil {
		return &AppError{err, "bad journal id", http.StatusBadRequest}
	}

	entries, err := database.GetJournalEntries(journalID)
	if err != nil {
		return &AppError{err, "error getting entries",
			http.StatusInternalServerError}
	}

	return renderJSON(w, entries, http.StatusOK)
}

func CreateJournalEntry(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	vars := mux.Vars(r)
	journalID, err := strconv.ParseInt(vars["journalID"], 10, 64)
	if err != nil {
		return &AppError{err, "bad journal id", http.StatusBadRequest}
	}

	defer r.Body.Close()
	entry, err := model.NewEntry(r.Body, journalID)
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

func Entry(w http.ResponseWriter, r *http.Request, u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return &AppError{err, "bad entry id",
			http.StatusBadRequest}
	}

	entry, err := database.GetEntry(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &AppError{err, "entry does not exist", http.StatusNotFound}
		}

		return &AppError{err, "error retrieving entry",
			http.StatusInternalServerError}
	}

	journal, err := database.GetJournal(entry.JournalID)
	if err != nil {
		return &AppError{err, "error retrieving entry",
			http.StatusInternalServerError}
	}

	if u.ID != journal.UserID {
		return &AppError{err, "not authorized", http.StatusUnauthorized}
	}

	return renderJSON(w, entry, http.StatusOK)
}
