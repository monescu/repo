package timetable

import "time"

//Course  contains all the information necessary about a given course in university
//Type is L, S, C
type Course struct {
	Year           string
	Specialisation string
	Group          string
	SemiGroup      string
	Name           string
	Type           string
	Location       string
	Teacher        string
	Day            time.Weekday
	Hours          string
	WeekType       string
}
