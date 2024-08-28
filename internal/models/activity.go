package models

type Activity struct {
	Title   string
	Start   string
	End     string
	Venue   string
	Host    string
	Status  int
	Remarks string
}

type Status map[int]string

var ActivityStatus = &Status{
	1: "To be conducted",
	2: "Conducted",
	3: "Rescheduled",
	4: "Postponed Indefinitely",
	5: "Canceled",
}
