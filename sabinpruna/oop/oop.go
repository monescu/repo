package oop

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type dog struct {
	name string
}

//makeNoise(azorel)
func makeNoise(d dog) {
	fmt.Printf("I am a dog, my name %v is  and i make noise.", d.name)
}

//func (r receiver) funcIdentifier returnedType {
//implementation
//}

//azorel.makeNoise()
func (d dog) makeNoise() {
	fmt.Printf("I am a dog, my name %v is  and i make noise.", d.name)
}

//----------------INTRO, VALIDATION, CTOR, GETTER SETTER---------------------------

//type typeName typeSpecification
type place struct {
	latitude, Longitude float64
	Name                string
}

func newPlace(latitude, longitude float64, name string) *place {
	return &place{latitude, longitude, name}
}

//  x := newPlace(0,0, "A")  -> *place
//  y := place{latitude: 1, Lon...} -> place
// x.Latitude() //fordny work
//  //work

//receiver value
// *T      *T
// T 		*T T

//getter
func (p *place) Latitude() float64 {
	return p.latitude
}

//setter
func (p *place) SetLatitude(latitude float64) {
	if latitude < 0 {
		p.latitude = 0
	} else {
		p.latitude = latitude
	}
}

func (p *place) String() string {
	return fmt.Sprintf("(%.3f°, %.3f°) %v", p.latitude, p.Longitude, p.Name)
}

//fmt.Println(place)
//fmt.println(place.string())

func (p *place) copy() *place {
	return &place{p.latitude, p.Longitude, p.Name}
}

//--------------------------------------------------------------------------

//----------ADDING METHODS---------------------------------------------------
type count int

func (count *count) increment()  { *count++ }
func (count *count) decrement()  { *count-- }
func (count count) isZero() bool { return count == 0 }

//Part is part
type Part struct {
	ID   int    // Named field (aggregation)
	Name string // Named field (aggregation)
}

//LowerCase method
func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

//UpperCase method
func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("«%d %q»", part.ID, part.Name)
}

//HasPrefix method
func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

//---------------------------------------------------------------------------

//---------------OVERRIDE---------------------------------------------------

type item struct {
	id       string  // Named field (aggregation)
	price    float64 // Named field (aggregation)
	quantity int     // Named field (aggregation)
}

func (item *item) cost() float64 {
	return item.price * float64(item.quantity)
}

type specialItem struct {
	item          // Anonymous field (embedding)
	catalogID int // Named field (aggregation)
}

type luxuryItem struct {
	item           // Anonymous field (embedding)
	markup float64 // Named field (aggregation)
}

/*
func (item *luxuryItem) cost() float64 {
	 // Needlessly verbose!
	 return item.Item.price * float64(item.Item.quantity) * item.markup }

func (item *luxuryItem) cost() float64 {
	// Needless duplication!
	return item.price * float64(item.quantity) * item.markup }
*/
func (item *luxuryItem) cost() float64 {
	// Ideal ✓
	return item.item.cost() * item.markup
}

//-------------------------------------------------------------------------

//--------------------INTERFACES-------------------------------------------

//Exchanger interface
type Exchanger interface {
	Exchange()
}

//StringPair struct
type StringPair struct{ first, second string }

//Exchange implementation
func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

//Point struct
type Point [2]int

//Exchange implementation
func (point *Point) Exchange() { point[0], point[1] = point[1], point[0] }

func (pair StringPair) String() string { return fmt.Sprintf("%q+%q", pair.first, pair.second) }

func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

//--------------------INTERFACES EMBEDDING-------------------------------------------

//LowerCaser interface
type LowerCaser interface{ LowerCase() }

//UpperCaser interface
type UpperCaser interface{ UpperCase() }

//LowerUpperCaser interface
type LowerUpperCaser interface {
	LowerCaser // As if we had written LowerCase()
	UpperCaser // As if we had written UpperCase()
}

//FixCaser interface
type FixCaser interface{ FixCase() }

//ChangeCaser interface
type ChangeCaser interface {
	LowerUpperCaser // As if we had written LowerCase(); UpperCase()
	FixCaser        // As if we had written FixCase()
}

//----------------------------------------------------------------------------------

//----------------------STRUCTS-----------------------------------------------------

//Person1 struct
type Person1 struct {
	Title     string   // Named field (aggregation)
	Forenames []string // Named field (aggregation)
	Surname   string   // Named field (aggregation)
}

//Author struct with aggregation (then anonymous field)
type Author struct {
	Person1           // Named field (aggregation)
	Title    []string // Named field (aggregation)
	YearBorn int      // Named field (aggregation)
}

//Tasks struct with anon func call
type Tasks struct {
	slice []string // Named field (aggregation)
	count          // Anonymous field (embedding)
}

//Add for gimmicky stuff
func (tasks *Tasks) Add(task string) {
	tasks.slice = append(tasks.slice, task)
	tasks.increment() // As if we had written: tasks.Count.Increment()
}

//----------------------------------------------------------------------------------

//------------------POLYMORPHISM----------------------------------------------------

//Income struct
type Income interface {
	calculate() int
	source() string
}

//FixedBilling struct
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

//TimeAndMaterial struct
type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

//-----------------------------------------------------------------------------------

func main() {
	{

		newYork := newPlace(40.716667, -74, "New York") // newYork is a *Place
		fmt.Println(newYork)
		baltimore := newYork.copy() // baltimore is a *Place
		baltimore.SetLatitude(newYork.Latitude() - 1.43333)
		baltimore.Name = "Baltimore"
		fmt.Println(baltimore)
	}

	{
		var c count //zero valued ,which for int is 0
		i := int(c)
		c.increment()
		j := int(c)
		c.decrement()
		k := int(c)
		fmt.Println(c, i, j, k, c.isZero())
	}

	{
		part := Part{5, "wrench"}
		part.UpperCase()
		part.ID += 11
		fmt.Println(part, part.HasPrefix("w"))

	}

	{
		special := specialItem{item{"Green", 3, 5}, 207}
		fmt.Println(special.id, special.price, special.quantity, special.catalogID)
		fmt.Println(special.cost())
	}

	{
		jekyll := StringPair{"Henry", "Jekyll"}
		hyde := StringPair{"Edward", "Hyde"}
		point := Point{5, -3}
		fmt.Println("Before: ", jekyll, hyde, point)
		jekyll.Exchange() // Treated as: (&jekyll).Exchange()
		hyde.Exchange()   // Treated as: (&hyde).Exchange()
		point.Exchange()  // Treated as: (&point).Exchange()
		fmt.Println("After #1:", jekyll, hyde, point)
		exchangeThese(&jekyll, &hyde, &point)
		fmt.Println("After #2:", jekyll, hyde, point)
	}

	{
		author := Author{
			Person1{"Mr", []string{"Robert", "Louis", "Balfour"}, "Stevenson"},
			[]string{"Kidnapped", "Treasure Island"},
			1850}
		fmt.Println(author)
		// author.Name.Title = ""
		// author.Name.Forenames = []string{"Oscar", "Fingal", "O'Flahertie", "Wills"}
		// author.Name.Surname = "Wilde"
		fmt.Println(author)
	}

	{
		project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
		project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
		project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
		incomeStreams := []Income{project1, project2, project3}
		calculateNetIncome(incomeStreams)
	}

	{

		//declare 3 persons
		sabin := Person{
			First: "sabin",
			Last:  "pruna",
			Age:   20,
		}
		andreea := Person{
			First: "andreea",
			Last:  "prandreeauna",
			Age:   22,
		}
		razvan := Person{
			First: "razvan",
			Last:  "razvan",
			Age:   22,
		}

		//create person slice
		people := []Person{sabin, andreea, razvan}

		//json marshal slice
		bs, err := json.Marshal(people)
		if err != nil {
			fmt.Println("didnt work")
		}

		fmt.Println(string(bs))

		//print string of bytes
	}

	{
		//et json , make it byte slice
		s := `[{"FirstName":"sabin","Last":"pruna","Age":25},{"FirstName":"andreea","Last":"prandreeauna","Age":22},{"FirstName":"razvan","Last":"razvan","Age":22}]`
		bs := []byte(s)
		fmt.Printf("%T\n", s)
		fmt.Printf("%T\n", bs)

		//json unmarhsal it &people
		var people []Person

		err := json.Unmarshal(bs, &people)
		if err != nil {
			fmt.Println("nope")
		}

		//iterate it
		for _, pers := range people {
			fmt.Println(pers)
		}

		//sort
		sort.Sort(ByAge(people))
		for _, pers := range people {
			fmt.Println(pers)
		}
	}
}
