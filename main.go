package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello There!")

	var name string = "Onur Belek"
	var status string = "working hard!"
	var weeklyGymCount int = 4

	fmt.Printf("%v is %v\nHe went to gym %v times this week\nHis favourite variable is %T\n", name, status, weeklyGymCount, weeklyGymCount)

	// PRINTING POINTERS
	fmt.Println(&name)

	// ARRAYS
	var randomArrNumbers = [50]int{10, 20, 30, 40, 999, 115, 666}

	// [50] is size of arr, int is data type, {} is the actual array you initialise and it can be empty

	// strongly typed like java - shows compiler errors like java too pretty nice

	fmt.Println(randomArrNumbers)

	var bookings [5]string
	// can add new elements to array with index location
	bookings[0] = name
	bookings[1] = "Donald Trump"
	bookings[2] = "Franklin Roosevelt"
	bookings[3] = "Herbert Hoover"
	bookings[4] = "Joe Biden"

	fmt.Println(len(bookings)) // len(arr) returns length of array like python

	// SLICES - abstraction of array - dynamic and powerful, index-based resized when needed - good for when you dont know what the length of the array will be

	var bookingSlice []string // don't need to define size

	bookingSlice = append(bookingSlice, "Donald Trump") // adding a value to the slice

	fmt.Println(bookingSlice)

	// FOR LOOPS - THIS IS AN INFINITE LOOP!
	// for {
	// 	fmt.Println(name)
	// }

	// for loops - iterate through array
	firstNames := []string{}

	for _, booking := range bookings { // range iterates over elements _ - is a blank variable to ignore variable you dont use
		var names = strings.Fields(booking) // .split
		firstNames = append(firstNames, names[0])
	}

	fmt.Println(firstNames)
}
