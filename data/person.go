package data

import (
	"math/rand"
)

type Person struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   uint8  `json:"age"`
}

var persons []Person = []Person{
	{
		ID:    1,
		Fname: "Travis",
		Lname: "Bickle",
		Age:   33,
	},
	{
		ID:    2,
		Fname: "Alice",
		Lname: "Cooper",
		Age:   75,
	},
}

func GetPersons() []Person {
	return persons
}

func GetPerson(id int) Person {
	return persons[0]
}

func (p *Person) AddPerson() {
	p.ID = rand.Intn(1000) + 1
	persons = append(persons, *p)
}
