package timetable

import (
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

//GetCourses returns a slice of all courses in the timetable
func GetCourses() []Course {
	unmerged, _ := xlsx.FileToSliceUnmerged("../Orar_sem_II_2018-2019_V4.xlsx")
	content := unmerged[0]

	courses := []Course{}

	for index, row := range content {
		if (row[0] != "1" && row[0] != "2" && row[0] != "3") || row[1] == "Codul Orarului: " {
			continue
		}

		year := row[0]
		spec := row[1]
		group := row[2]
		semi := row[3]

		for i := 4; i < len(row); i++ {
			name, courseType, location, teacher := parseCourseCell(row[i])

			if name != "" {
				courses = append(courses, Course{
					Day:            getDay(i),
					Group:          group,
					Hours:          getHours(i),
					Location:       location,
					Name:           name,
					SemiGroup:      semi,
					Specialisation: spec,
					Teacher:        teacher,
					Type:           courseType,
					WeekType:       getWeekType(index),
					Year:           year})
			}
		}
	}

	return courses
}

func search(courses []Course, day time.Weekday, discipline string, location string, year string, spec string, teacher string, courseType string, semi string) []Course {
	return []Course{}
}

//----------------------------------PRIVATES--------------------------------------------

func getWeekType(index int) string {
	if index%2 == 0 {
		return "Even"
	}

	return "Odd"
}

func getHours(index int) string {

	dayIndex := (index - 4) % 7

	switch dayIndex {
	case 0:
		return "8,00-9,50"
	case 1:
		return "10,00-11,50"
	case 2:
		return "12,00-13,50"
	case 3:
		return "14,00-15,50"
	case 4:
		return "16,00-17,50"
	case 5:
		return "18,00-19,50"
	case 6:
		return "20,00-21,50"
	}

	return "Unknown"
}

func getDay(index int) time.Weekday {
	switch {
	case index >= 4 && index <= 10:
		return time.Monday
	case index >= 11 && index <= 17:
		return time.Tuesday
	case index >= 18 && index <= 24:
		return time.Wednesday
	case index >= 25 && index <= 31:
		return time.Thursday
	case index >= 32 && index <= 38:
		return time.Thursday
	default:
		return time.Sunday
	}
}

func parseCourseCell(row string) (string, string, string, string) {
	if row == "" {
		return "", "", "", ""
	}
	splitted := strings.Split(row, ",")

	if len(splitted) != 4 {
		return "", "", "", ""
	}

	return splitted[0], splitted[1], splitted[2], splitted[3]
}

//--------------------------------------------------------------------------------------
