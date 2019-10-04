package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Inside reflect test")

	//! type identification
	//! Typeof
	//! Valueof
	var i []int
	i = []int{1, 2, 3}
	fmt.Println("type of i is ", reflect.TypeOf(i).Kind())
	fmt.Println("value of i is ", reflect.ValueOf(i))

	//! append
	n := []int {1,2,3}
	t := reflect.Append(reflect.ValueOf(n), reflect.ValueOf(5))
	fmt.Println(t)
}