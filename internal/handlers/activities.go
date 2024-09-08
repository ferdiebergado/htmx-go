package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ferdiebergado/htmx-go/internal/models"
	"github.com/ferdiebergado/htmx-go/internal/services"
	"github.com/ferdiebergado/htmx-go/internal/view"
)

type ActivityHandler struct {
	Repository models.ActivityRepository
	Session    *services.SessionManager
}

func (h *ActivityHandler) ListActivities(w http.ResponseWriter, r *http.Request) {

	var activities []interface{}

	result, err := h.Repository.GetAllActivities(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert each element in result to interface{} and append to activities
	for _, activity := range result {
		activities = append(activities, activity)
	}

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   map[string]interface{}{"activities": activities},
	}

	view.Render(w, r, "activities.html", data)
}

func (h *ActivityHandler) ShowActivity(w http.ResponseWriter, r *http.Request) {
	activity, err := extractActivity(r, h.Repository)

	if err != nil {
		log.Println("Invalid id:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if activity.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   map[string]interface{}{"activity": activity},
	}

	view.Render(w, r, "activity.html", data)
}

func (h *ActivityHandler) ShowActivityForm(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "PTMS - New Activity",
		Icon:   "calendar-o",
		Header: "Activities",
	}

	view.Render(w, r, "activity-form.html", data)
}

func (h *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println("Error parsing form data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activityStatus, err := strconv.Atoi(r.FormValue("status"))

	if err != nil {
		log.Println("Invalid status", err)
		w.WriteHeader(http.StatusBadRequest)
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

	_, err = h.Repository.CreateActivity(r.Context(), activity)

	if err != nil {
		log.Println("failed creating activity", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (h *ActivityHandler) ShowActivityEditForm(w http.ResponseWriter, r *http.Request) {
	activity, err := extractActivity(r, h.Repository)

	if err != nil {
		log.Println("Invalid id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   map[string]interface{}{"activity": activity},
	}

	view.Render(w, r, "activity-form.html", data)
}

func (h *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println("Invalid id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activityStatus, err := strconv.Atoi(r.FormValue("status"))

	if err != nil {
		log.Println("Invalid status", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activity, err := extractActivity(r, h.Repository)

	if err != nil {
		log.Println("Invalid id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activity.Title = r.FormValue("title")
	activity.Start = r.FormValue("start_date")
	activity.End = r.FormValue("end_date")
	activity.Venue = r.FormValue("venue")
	activity.Host = r.FormValue("host")
	activity.Status = activityStatus
	activity.Remarks = r.FormValue("remarks")

	err = h.Repository.UpdateActivity(r.Context(), activity)

	if err != nil {
		log.Println("failed to update activity", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (h *ActivityHandler) DestroyActivity(w http.ResponseWriter, r *http.Request) {
	activity, err := extractActivity(r, h.Repository)

	if err != nil {
		log.Println("Invalid id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if activity.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.Repository.DeleteActivity(r.Context(), activity.ID)

	if err != nil {
		log.Println("failed to delete activity", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func extractActivity(r *http.Request, repo models.ActivityRepository) (*models.Activity, error) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		return &models.Activity{}, err
	}

	log.Println("ACTIVITY ID:", id)

	activity, err := repo.GetActivity(r.Context(), id)

	if err != nil {
		return &models.Activity{}, err
	}

	return activity, nil
}
