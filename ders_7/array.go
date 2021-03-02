package main

import (
	"fmt"
)

func init() {
	fmt.Println("Array tutorials...")
}

func main(){
	//var numbers3 []int = []int{2,3,4,5,6}

	var numbers1 = []int{1,2,3,4}
	fmt.Printf("%v\n", numbers1)

	numbers2  := numbers1

	numbers2[0] = 21
	fmt.Printf("%v\n", numbers1)
	fmt.Printf("%v\n", numbers2)


	var numbers3 []int = numbers1
	numbers3[0] = 43
	fmt.Printf("%v\n", numbers1)
	fmt.Printf("%v\n", numbers3)


	var cities1 []string = []string{"Adana", "İzmir", "İstanbul", "Diyarbakır"}
	cities2 :=  [4]string{}

	//copy(cities2, cities1)

	cities2[0] = "Çorum"

	fmt.Printf("%v\n", cities1)
	fmt.Printf("%v\n", cities2)


	var ages = []int{21,35,43}
	test(ages)
	fmt.Printf("%v\n", ages)



	var age int = 21
	test2(age);

	fmt.Printf("%d\n", age);


	var a [5]int
	fmt.Printf("%v\n", a)

	var x interface{}
	x = 21
	x = 35.8
	x = "enes"
	x = []int{21,35,43}
	fmt.Println(x)

	a1 := [10]int{}

	fmt.Println(a1)


	values := []int{1, 2, 3, 4}
	//values[4] = 21
	values = append(values, 43)
	values = append(values, 35)
	fmt.Println(values)
}


//call by value
func test2(age int){
	age = 200
}

//call by reference  &
func test(numbers []int){
	numbers[0] = 20
}