package main

import (
	"fmt"
	"reflect"
)

func main() {

	myArray := [5]int{18, 25, 30, 42, 50}
	fmt.Println(myArray)

	/*
		Here we craete a slice that spans across the whole existing array.
		Which means that this slice can also control whats stored in each
		array item
	*/
	mySlice := myArray[:]
	fmt.Println(mySlice)

	/*
		It works a bit like hard-links. So modifying the slice will end up modifying
		the array, and vice versa
	*/
	myArray[0] = 1
	mySlice[1] = 2
	fmt.Println(myArray) // prints [1 2 30 42 50]
	fmt.Println(mySlice) // prints [1 2 30 42 50]

	/*
	   Creating a slice seems to be a lot work. You first create an array, then a slice from that array.
	   It also means that your limited to the maximum size of the array which in the above example was 5.
	   Here's another more short-hand and powerful way to create a slice

	   The empty square bracket here tells golang that this is a slice not an array.
	   Behind the scenes golang will create a 5 item array to store the slices data.
	*/

	my2ndSlice := []int{4, 5, 3, 2, 10}
	fmt.Println(reflect.TypeOf(my2ndSlice))
	fmt.Println(my2ndSlice)

	/*
		Here, golang creates a new, bigger, array behind the scenes, copies the existing slice's
		array's content, and append the new additional values. It then deletes the old hidden array.
	*/
	my2ndSlice = append(my2ndSlice, 25, 45, 30) // prints [4 5 3 2 10 25 45 30]
	fmt.Println(my2ndSlice)

	/*
	   Here's how to create a slice using a subset of an existing slice
	*/

	my3rdSlice := my2ndSlice[3:5] // prints [2 10]
	fmt.Println(my3rdSlice)
}
