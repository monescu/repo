package oop

//Person struct
type Person struct {
	First string `json:"FirstName"`
	Last  string //`json:"-"`
	Age   int
}

//ByAge used for sorting
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
