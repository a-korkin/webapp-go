package data

import (
	"log"
)

type Person struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   uint8  `json:"age"`
}

func GetPersons(appState *AppState) []Person {
	if appState.Db == nil {
		log.Fatalf("db not initialized")
	}
	rows, err := appState.Db.Query(`
select id, lname, fname, age 
from public.person`)
	if err != nil {
		log.Fatalf("failed to get persons: %s", err)
	}

	persons := []Person{}
	for rows.Next() {
		pers := Person{}
		err := rows.Scan(&pers.ID, &pers.Lname, &pers.Fname, &pers.Age)
		if err != nil {
			log.Fatalf("failed to prepare person: %s", err)
		}
		persons = append(persons, pers)
	}
	return persons
}

func GetPerson(id int, appState *AppState) (*Person, error) {
	row, err := appState.Db.Query(`
select id, lname, fname, age 
from public.person 
where id = $1`, id)
	if err != nil {
		log.Fatalf("failed to get person: %s", err)
		return nil, err
	}
	pers := Person{}
	if row.Next() {
		if err = row.Scan(&pers.ID, &pers.Lname, &pers.Fname, &pers.Age); err != nil {
			log.Fatalf("failed to prepare person: %s", err)
		}
	}
	return &pers, nil
}

func (p *Person) AddPerson(appState *AppState) {
	res, err := appState.Db.Query(`
insert into public.person(lname, fname, age)
values($1, $2, $3)
returning id`, p.Lname, p.Fname, p.Age)
	if res.Next() {
		if err = res.Scan(&p.ID); err != nil {
			log.Fatalf("failed to get id: %s", err)
		}
	}
	if err != nil {
		log.Fatalf("failed to create person: %s", err)
	}
}

func DeletePerson(id int, appState *AppState) {
	_, err := appState.Db.Exec(`
delete from public.person where id = $1`, id)
	if err != nil {
		log.Fatalf("failed to delete person: %s", err)
	}
}
