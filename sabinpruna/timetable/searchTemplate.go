package timetable

import (
	"sort"
	"time"
)

//SearchTemplate is used to populate the form's dropdown values
type SearchTemplate struct {
	Days            []time.Weekday
	Disciplines     []string
	Locations       []string
	Years           []string
	Specialisations []string
	Teachers        []string
	CourseTypes     []string
	Groups          []string
	SemiGroups      []string
	Hours           []string
	MatchedCourses  []Course
	IsPost          bool
}

//GetSearchTemplate returns a new SearchTemplate with distinct values for all dropdown fields after reading the courses from the timetable
func GetSearchTemplate(courses []Course) SearchTemplate {

	disciplines, locations, years, specs, teachers, groups := getFields(courses)

	return SearchTemplate{
		Days:            []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
		Disciplines:     disciplines,
		Locations:       locations,
		Years:           years,
		Specialisations: specs,
		Teachers:        teachers,
		CourseTypes:     []string{"*", "L", "S", "C"},
		Groups:          groups,
		SemiGroups:      []string{"*", "A", "B"},
		Hours:           []string{"*", "8,00-9,50", "10,00-11,50", "12,00-13,50", "14,00-15,50", "16,00-17,50", "18,00-19,50", "20,00-21,50"},
		IsPost:          false,
	}
}

//Search returns a list of courses that match the criteria
func Search(courses []Course, teacher string, day string, discipline string, year string, specialisation string, courseType string, group string, semiGroup string, hours string) []Course {
	matchedCourses := []Course{}

	//string(course.Day) == day &&
	//course.Teacher == teacher &&
	//course.Name == discipline &&
	//course.Year == year &&
	//course.Specialisation == specialisation &&
	//course.Type == courseType &&
	//course.Group == group &&
	//course.SemiGroup == semiGroup

	for _, course := range courses {
		ok := true

		if toStringWeekday(course.Day) != day {
			ok = false
		}
		if teacher != "*" && course.Teacher != teacher {
			ok = false
		}
		if discipline != "*" && course.Name != discipline {
			ok = false
		}
		if year != "*" && course.Year != year {
			ok = false
		}
		if specialisation != "*" && course.Specialisation != specialisation {
			ok = false
		}
		if courseType != "*" && course.Type != courseType {
			ok = false
		}
		if group != "*" && course.Group != group {
			ok = false
		}
		if semiGroup != "*" && course.SemiGroup != semiGroup {
			ok = false
		}

		if hours != "*" && course.Hours != hours {
			ok = false
		}

		if ok {
			matchedCourses = append(matchedCourses, course)
		}
	}

	return matchedCourses
}

//-----------------------------PRIVATES------------------------------------

func removeSliceDuplicates(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func getFields(courses []Course) ([]string, []string, []string, []string, []string, []string) {
	disciplines := []string{"*"}
	locations := []string{"*"}
	years := []string{"*"}
	specs := []string{"*"}
	teachers := []string{}
	groups := []string{"*"}

	for _, course := range courses {
		disciplines = append(disciplines, course.Name)
		locations = append(locations, course.Location)
		years = append(years, course.Year)
		specs = append(specs, course.Specialisation)
		teachers = append(teachers, course.Teacher)
		groups = append(groups, course.Group)
	}

	disciplines = removeSliceDuplicates(disciplines)
	locations = removeSliceDuplicates(locations)
	years = removeSliceDuplicates(years)
	specs = removeSliceDuplicates(specs)
	teachers = removeSliceDuplicates(teachers)
	groups = removeSliceDuplicates(groups)

	sort.Strings(disciplines)
	sort.Strings(locations)
	sort.Strings(years)
	sort.Strings(specs)
	sort.Strings(teachers)
	sort.Strings(groups)

	//only because someone added an extra space and there are 2 similar values that are different by a space
	locations = locations[1:]

	teachers = append([]string{"*"}, teachers...)

	return disciplines, locations, years, specs, teachers, groups
}

func toStringWeekday(day time.Weekday) string {
	// declare an array of strings
	// ... operator counts how many
	// items in the array (7)
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}
	// → `day`: It's one of the
	// values of Weekday constants.
	// If the constant is Sunday,
	// then day is 0.
	//
	// prevent panicking in case of
	// `day` is out of range of Weekday
	if day < time.Sunday || day > time.Saturday {
		return "Unknown"
	}
	// return the name of a Weekday
	// constant from the names array
	// above.
	return names[day]
}

//-------------------------------------------------------------------------
