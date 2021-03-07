package main

import "fmt"

func init() {
	fmt.Println("Defer Tutorials...")
}

func main(){
	//defer fmt.Println("World")
	//defer fmt.Println("Golang Tutorials...")
	fmt.Println("Hello")

	/*
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	*/
	//defer test1() 
	//test2()
	//fmt.Println("Panic before")
	//panic("Not working!!")
	//fmt.Println("Panic after")

	//defer defFunc()
	//panic("Not working!!")

	val := c()
	fmt.Println(val)
	ageControl(-23)
	
}

func ageControl(age int) int{

	if age < 0{
		panic("Age must not negative")
	}

	return age
}

func test3() (name string, number int){
	name = "ali"
	number = 5
	return 
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}

func defFunc() {
    fmt.Println("This is a deferred function")
}

func test1(){
	fmt.Println("Test1")
}

func test2(){
	fmt.Println("Test2")
}