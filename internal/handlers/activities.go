package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/models"
	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func HandleActivities(w http.ResponseWriter, _ *http.Request) {

	var activities []interface{}

	activity1 := models.Activity{
		Title:  "Monitoring of the Great Red Spot for Atmospheric Anomalies",
		Start:  "9/15/2024",
		End:    "9/19/2024",
		Venue:  "Jupiter",
		Host:   "Thanos",
		Status: "To be conducted",
	}

	activity2 := models.Activity{
		Title:  "Training on Interstellar Warfare using Magnetars",
		Start:  "11/5/2024",
		End:    "11/29/2024",
		Venue:  "Andromeda Galaxy",
		Host:   "Galactus",
		Status: "To be conducted",
	}

	activities = append(activities, activity1, activity2)

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
		Data:   activities,
	}

	utils.Render(w, "activities.html", data)

}

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	data := &PageData{
		Title:  "PTMS - New Activity",
		Icon:   "calendar-o",
		Header: "Activities",
	}

	utils.Render(w, "create-activity.html", data)
}
