package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ferdiebergado/htmx-go/internal/db"
)

type Activity struct {
	Model
	Title   string
	Start   string
	End     string
	Venue   string
	Host    string
	Status  int
	Remarks string
}

type ActivityRepository interface {
	GetAllActivities(ctx context.Context) ([]*Activity, error)
	GetActivity(ctx context.Context, id int) (*Activity, error)
	CreateActivity(ctx context.Context, activity *Activity) (int, error)
	UpdateActivity(ctx context.Context, activity *Activity) error
	DeleteActivity(ctx context.Context, id int) error
}

type activityRepository struct {
	DB db.Database
}

var ActivityStatus = map[int]string{
	1: "To be conducted",
	2: "Conducted",
	3: "Rescheduled",
	4: "Postponed Indefinitely",
	5: "Cancelled",
}

func NewActivityRepository(db db.Database) *activityRepository {
	return &activityRepository{DB: db}
}

func (r *activityRepository) GetAllActivities(ctx context.Context) ([]*Activity, error) {

	const query = `
	SELECT id, created_at, updated_at, title, start_date, end_date,  venue, host, status, remarks
	FROM activities
	WHERE deleted_at IS NULL
	ORDER BY start_date DESC
	`
	var activities []*Activity

	rows, err := r.DB.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	err = db.MarshalRowsToStructs(rows, &activities)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (r *activityRepository) GetActivity(ctx context.Context, id int) (*Activity, error) {
	const query = `
	SELECT id, created_at, updated_at, title, start_date, end_date, venue, host, status, remarks 
	FROM activities
	WHERE deleted_at IS NULL AND id = $1
	`

	var activity Activity

	row := r.DB.QueryRow(ctx, query, id)

	err := db.MarshalRowToStruct(row, &activity)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (r *activityRepository) CreateActivity(ctx context.Context, activity *Activity) (int, error) {
	const query = `
	INSERT INTO activities
	(title, start_date, end_date, venue, host, status, remarks) 
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id`

	var id int

	row := r.DB.QueryRow(ctx, query, activity.Title, activity.Start, activity.End, activity.Venue, activity.Host, activity.Status, activity.Remarks)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *activityRepository) UpdateActivity(ctx context.Context, activity *Activity) error {
	const query = `
	UPDATE activities 
	SET title = $1, start_date = $2, end_date = $3, venue = $4, host = $5, status = $6, remarks = $7 
	WHERE id = $8`

	_, err := r.DB.Exec(ctx, query, activity.Title, activity.Start, activity.End, activity.Venue, activity.Host, activity.Status, activity.Remarks, activity.ID)

	return err
}

func (r *activityRepository) DeleteActivity(ctx context.Context, id int) error {
	const query = `UPDATE activities SET deleted_at = NOW() WHERE id = $1`

	_, err := r.DB.Exec(ctx, query, id)

	return err
}

func (a *Activity) ParseStatus(id int) string {
	return fmt.Sprint(ActivityStatus[id])
}

func (a *Activity) ParseDate(timeString string) string {

	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		log.Println("Error parsing time:", err)
		return timeString
	}

	return fmt.Sprint(t.Format(time.DateOnly))
}
