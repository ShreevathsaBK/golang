package main

import "fmt"

type person struct{
	name string
	age int
}

func (p *person) printPerson() {
	fmt.Printf("Name is %v and age is %v\n", p.name, p.age)
}

func newPerson(name string, age int) person{
	p := person{name : name, age : age}
	return p
}

func main(){
	p1 := newPerson("shreevathsa", 20)

	fmt.Println("p1 :", p1)
	fmt.Println("p1.name :", p1.name)
	fmt.Println("p1.age :", p1.age)

	p1.printPerson()
}