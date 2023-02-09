package main

import (
	"fmt"
)

type Person struct {
	id   int
	name string
}

func main() {

	fmt.Println("hello man")

	slices := []int{1, 2, 3, 4, 5}
	fmt.Printf("slices 3: %v\n", slices)
	slices = append(slices, 6)
	slices = RemoveIndex(slices, 3)
	fmt.Printf("slices 4: %v\n", slices)

	// _ = append(s[:3], s[4:]...)

	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
	all = RemoveIndex(all, 5)
	fmt.Println(all) //[0 1 2 3 4 6 7 8 9]

	people := []Person{}

	people = append(people, Person{1, "john"})
	people = append(people, Person{2, "kim"})
	people = append(people, Person{3, "Annie"})
	people = append(people, Person{4, "Jada"})

	fmt.Printf("\n Create %v\n", people)

	people = RemoveByID(people, 3)

	fmt.Printf("\n Delete %v\n", people)

	people = UpdatePerson(people, 2, Person{2, "Hannah"})

	fmt.Printf("\n Update %v\n", people)

}

func UpdatePerson(s []Person, id int, person Person) []Person {

	for i, v := range s {

		if v.id == id {

			s[i] = person
			return s
		}
	}

	return []Person{{-1, "no such id"}}
}

func RemoveByID(s []Person, id int) []Person {

	for i, v := range s {

		if v.id == id {
			return RemovePersonIndex(s, i)
		}
	}

	return []Person{{-1, "no such id"}}

}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemovePersonIndex(s []Person, index int) []Person {
	return append(s[:index], s[index+1:]...)
}
