package main

import "fmt"

type Gender uint

const (
	Male Gender = iota + 1
	Female
)

var genders = []string{
	"Male",
	"Female",
}

func (g Gender) String() string {
	return genders[g-1]
}

type Person struct {
	Gender Gender
	Age    int
	Name   string
}

var persons = []Person{
	{Age: 25, Name: "John Doe"},
	{Age: 25, Name: "Adam Smith"},
}

// Answer from Stackoverflow
// http://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func contains(people []Person, p Person) bool {
	for _, per := range people {
		if per.Name == p.Name {
			return true
		}
	}
	return false
}

func main() {
	p := Person{Name: "Adam Smith", Gender: Female}
	if contains(persons, p) {
		fmt.Printf("%#v - %s\n is in persons", p, p.Gender.String())
	} else {
		fmt.Printf("%#v - %s\n is not in persons", p, p.Gender.String())
	}
}
