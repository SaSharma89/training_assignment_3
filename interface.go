/*
 * Implement interface with N number of methods and implement
 * only N-1 method and call all methods using type associate with them
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

//func (p person)familyInfo(){
//	fmt.Println("person's family info")
//}

type person struct {}

func main () {
	fmt.Println("Inside interface first assignment")

	per := person{}

	per.basicInfo()
	per.contactInfo()

	//per.familyInfo()
	//
	//var test intf
	//test = per
	//
	//test.contactInfo()
}

/*
-> if type has not implemented all methods than its not able to call all.
-> If all methods are not able to implemented than not able to assign person type into intf type
 */