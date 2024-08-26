package models

type Travel struct {
	Activity   string
	StartDate  string
	EndDate    string
	Venue      string
	Personnels []Personnel
	Status     uint8
}

func (t *Travel) NewTravel(newTravel Travel) *Travel {
	return new(Travel)
}
