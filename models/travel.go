package models

type Travel struct {
	Activity   string
	StartDate  string
	EndDate    string
	Venue      string
	Personnels []Personnel
}

func (t *Travel) New(newTravel Travel) {

}
