package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zachlatta/angelhack/database"
	"github.com/zachlatta/angelhack/model"
)

func Journals(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	journals, err := database.GetJournals(u.ID)
	if err != nil {
		return &AppError{err, "error getting journals",
			http.StatusInternalServerError}
	}

	return renderJSON(w, journals, http.StatusOK)
}

func CreateJournal(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	defer r.Body.Close()
	journal, err := model.NewJournal(r.Body, u.ID)
	if err != nil {
		return &AppError{err, err.Error(), http.StatusBadRequest}
	}

	err = database.SaveJournal(journal)
	if err != nil {
		return &AppError{err, "error saving journal",
			http.StatusInternalServerError}
	}

	return renderJSON(w, journal, http.StatusOK)
}

func Journal(w http.ResponseWriter, r *http.Request, u *model.User) *AppError {
	if u == nil {
		return &AppError{ErrUnauthorized, "not authorized",
			http.StatusUnauthorized}
	}

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return &AppError{err, "bad journal id", http.StatusBadRequest}
	}

	journal, err := database.GetJournal(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &AppError{err, "journal does not exist", http.StatusNotFound}
		}

		return &AppError{err, "error retrieving journal",
			http.StatusInternalServerError}
	}

	if journal.UserID != u.ID {
		return &AppError{err, "not authorized", http.StatusUnauthorized}
	}

	return renderJSON(w, journal, http.StatusOK)
}
