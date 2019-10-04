package main

import "fmt"

type intf interface {
	display()
}

type person struct {
	name string
	age  int
}

//func (p *person) display() {
//	fmt.Println("Inside pointer display : name ", p.name, " age ", p.age)
//	p.age += 2
//}


func (p person) display() {
	fmt.Println("Inside normal display : name ", p.name, " age ", p.age)
	p.age += 2
}

func main() {
	fmt.Println("First program of 3 assignment")
	per := person{ "SANDEEP", 30 }
	per.display()

	fmt.Println("Inside main : name ", per.name, " age ", per.age)
}

/*
So from above analysis,
-> while using pointer method its passing value by refrence so any changes in parameter will reflects in main
-> while using object method its create another object so any changes in value will not reflects in main
 */