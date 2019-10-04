/*
 * Implement interface with N number of methods and implement
 * all method and call all methods using type associate with them.
 * Extend program with multiple types.
 */

package main
import "fmt"


type intf interface {
	basicInfo()
	contactInfo()
	familyInfo()
}


func (p person)basicInfo(){
	fmt.Println("person's basic info")
}

func (p person)contactInfo(){
	fmt.Println("person's contact info")
}

func (p person)familyInfo(){
	fmt.Println("person's family info")
}

type person struct {}

func (e employee)basicInfo(){
	fmt.Println("employee's basic info")
}

func (e employee)contactInfo(){
	fmt.Println("employee's contact info")
}

func (e employee)familyInfo(){
	fmt.Println("employee's family info")
}


type employee struct {}

func main () {
	fmt.Println("Inside interface second assignment")
	var in int
	var test intf

	fmt.Println("Enter 1 for person details and 2 for employee details")
	fmt.Scan(&in)

	switch in {
	case 1 : {
		test = person{}
	}
	case 2 : {
		test = employee{}
	}
	}

	test.basicInfo()
	test.contactInfo()
	test.familyInfo()
}