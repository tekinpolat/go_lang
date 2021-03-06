package main

import "fmt"

func init() {
	fmt.Println("Slice Tutorials...")
}

func main() {
	//array fix size
	//slice resize
	var numberArr = [5]int{1, 2, 3, 4, 6} //array
	var numberSlice []int                 //slice

	fmt.Println(numberArr)
	fmt.Println(numberSlice)

	//Slices are reference-types while arrays are value-type.

	var stringSlice = []string{"This", "is", "a", "string", "slice"}  //slice
	stringSlice = append(stringSlice, "test") //add
	// ?stringSlice[6] = "test2"  wrong

	fmt.Println(stringSlice)

	var odds = [8]int{1, 3, 5, 7, 9, 11, 13, 15}
	slice1 := odds[2:]
	slice1[0] = 21

	fmt.Println(odds)
	fmt.Println(slice1)


	// !slice
	numbers1 := make([]int, 5)
	numbers1 = append(numbers1, 43)
	fmt.Println(numbers1)

	numbers2 := make([]int, 5, 5)   //??
	numbers2 = append(numbers2, 21)
	fmt.Println(numbers2)


	
	//?numbers3 := []int{}
	//!var numbers4 []int

	numbers3 := []int{1, 2, 3, 4}
	//numbers3[1:3] = 21 wrong
	_ = append(numbers3, 21)
	fmt.Printf("%v\n", numbers3)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//s = append(s, 9, 10)
	
	arr := []int{21,35,43}
	s = append(s, arr...)
	fmt.Println(s)
	/*
	for _, val := range arr{
		s = append(s, val)
	}
	*/

	arr1 := [] int{1,2,4,5}

	for index, val := range arr1{
		fmt.Printf("%d %p %T\n", val, &arr1[index], val)
	}

}
