package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ferdiebergado/htmx-go/internal/models"
	"github.com/ferdiebergado/htmx-go/internal/services"
	"github.com/ferdiebergado/htmx-go/internal/utils"
)

type ActivityHandler struct {
	Repository models.ActivityRepository
	Session    *services.SessionManager
}

func (h *ActivityHandler) ListActivities(w http.ResponseWriter, r *http.Request) {

	var activities []interface{}

	result, err := h.Repository.GetAllActivities(r.Context())

	if err != nil {
		panic(err)
	}

	// Convert each element in result to interface{} and append to activities
	for _, activity := range result {
		activities = append(activities, activity)
	}

	// const query = `SELECT id, title, start_date, end_date, venue, host, status, remarks FROM activities`

	// if err != nil {
	// 	panic(err)
	// }

	// defer rows.Close()

	// var activities []interface{}

	// for rows.Next() {

	// 	var activity models.Activity

	// 	if err := rows.Scan(
	// 		&activity.ID,
	// 		&activity.Title,
	// 		&activity.Start,
	// 		&activity.End,
	// 		&activity.Venue,
	// 		&activity.Host,
	// 		&activity.Status,
	// 		&activity.Remarks,
	// 	); err != nil {
	// 		panic(err)
	// 	}

	// 	activities = append(activities, activity)
	// }

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   activities,
	}

	utils.Render(w, "activities.html", data)
}

func (h *ActivityHandler) ShowActivity(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		panic(err)
	}

	activity, err := h.Repository.GetActivity(r.Context(), id)

	if err != nil {
		panic(err)
	}

	// const query = `
	// SELECT id, title, start_date, end_date, venue, host, status, remarks, created_at, updated_at
	// FROM activities
	// WHERE id = $1`

	// id := r.PathValue("id")

	// rows, err := h.DB.Query(r.Context(), query, id)

	// if err != nil {
	// 	panic(err)
	// }

	// defer rows.Close()

	// var activity models.Activity

	// for rows.Next() {

	// 	if err := rows.Scan(
	// 		&activity.ID,
	// 		&activity.Title,
	// 		&activity.Start,
	// 		&activity.End,
	// 		&activity.Venue,
	// 		&activity.Host,
	// 		&activity.Status,
	// 		&activity.Remarks,
	// 		&activity.CreatedAt,
	// 		&activity.UpdatedAt,
	// 	); err != nil {
	// 		panic(err)
	// 	}
	// }

	fmt.Printf("activity: %v\n", activity)

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   []interface{}{activity},
	}

	utils.Render(w, "activity.html", data)
}

func (h *ActivityHandler) ShowActivityForm(w http.ResponseWriter, r *http.Request) {
	var payloads []interface{}

	session := r.Context().Value(services.SessionKey{}).(*services.Session)
	csrfToken, _ := h.Session.SetCSRFToken(session)

	log.Println("csrfToken: " + csrfToken)

	payloads = append(payloads, csrfToken)

	data := &PageData{
		Title:  "PTMS - New Activity",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   payloads,
	}

	utils.Render(w, "create-activity.html", data)
}

func (h *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	// const query = `
	// INSERT INTO activities
	// (title, start_date, end_date, venue, host, status,remarks)
	// VALUES ($1,$2,$3,$4,$5,$6,$7)`

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	activityStatus, err := strconv.Atoi(r.FormValue("status"))

	if err != nil {
		http.Error(w, "Invalid status input", http.StatusBadRequest)
		return
	}

	activity := &models.Activity{
		Title:   r.FormValue("title"),
		Start:   r.FormValue("start_date"),
		End:     r.FormValue("end_date"),
		Venue:   r.FormValue("venue"),
		Host:    r.FormValue("host"),
		Status:  activityStatus,
		Remarks: r.FormValue("remarks"),
	}

	// title := r.FormValue("title")
	// start := r.FormValue("start_date")
	// end := r.FormValue("end_date")
	// venue := r.FormValue("venue")
	// host := r.FormValue("host")
	// status := activityStatus
	// remarks := r.FormValue("remarks")

	_, err = h.Repository.CreateActivity(r.Context(), activity)

	if err != nil {
		panic(err)
	}

	// defer rows.Close()

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}
