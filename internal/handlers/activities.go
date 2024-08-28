package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ferdiebergado/htmx-go/internal/db"
	"github.com/ferdiebergado/htmx-go/internal/models"
	"github.com/ferdiebergado/htmx-go/internal/utils"
)

type ActivityHandler struct {
	DB db.Database
}

func (h *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	activityStatus, err := strconv.Atoi(r.FormValue("status"))

	if err != nil {
		http.Error(w, "Invalid age input", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	start := r.FormValue("start_date")
	end := r.FormValue("end_date")
	venue := r.FormValue("venue")
	host := r.FormValue("host")
	status := activityStatus
	remarks := r.FormValue("remarks")

	query := `INSERT INTO activities 
	(title, start_date, end_date, venue, host, status,remarks) 
	VALUES ($1,$2,$3,$4,$5,$6,$7)`

	rows, err := h.DB.Query(r.Context(), query, title, start, end, venue, host, status, remarks)

	if err != nil {
		log.Println(fmt.Errorf("an error occured: %v", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "/activities", http.StatusFound)
}

func (h *ActivityHandler) ListActivities(w http.ResponseWriter, r *http.Request) {

	sql := `SELECT id, title, start_date, end_date, venue, host, status, remarks FROM activities`

	rows, err := h.DB.Query(r.Context(), sql)

	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	defer rows.Close()

	var activities []interface{}

	for rows.Next() {

		var activity models.Activity

		if err := rows.Scan(
			&activity.ID,
			&activity.Title,
			&activity.Start,
			&activity.End,
			&activity.Venue,
			&activity.Host,
			&activity.Status,
			&activity.Remarks,
		); err != nil {
			utils.HandleHTTPError(w, err)
			return
		}

		activities = append(activities, activity)
	}

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   activities,
	}

	utils.Render(w, "activities.html", data)
}

func (h *ActivityHandler) ShowActivityForm(w http.ResponseWriter, r *http.Request) {
	data := &PageData{
		Title:  "PTMS - New Activity",
		Icon:   "calendar-o",
		Header: "Activities",
	}

	utils.Render(w, "create-activity.html", data)
}
