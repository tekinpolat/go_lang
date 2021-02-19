package main

import "fmt"

func main() {
	/*
		Arithmetic Operators               + - * / %
		Comparison (Relational) Operators  <= < > >=
		Assignment Operators               = %= /= //= -= += *= **=
		Logical Operators                  && || !                   and or not
		Bitwise Operators                  & |
		Membership Operators
		Identity Operators
	*/

	//0000 1101
	//0000 0110

	var number1 float32 = 3.5 //32 bit
	var number2 float32 = 5.6
	//var number3 int = 34

	var add float32 = number1 + number2

	fmt.Println("Add =>", add)
	fmt.Println(number1 == number2)

	fmt.Println(true && false)

	fmt.Println(!true)

	var number3 int = 21
	var number4 int = 43

	number3 += number4 // number3 = number3 + number4

	var number5 int = 5
	var number6 int = number5 << 3 //  0101 => 5  1010

	fmt.Println(number6)

	fmt.Println(number3 | number4)

	var name string = "enes"

	surname := "Karacaer"

	var age int
	age = 35

	fmt.Printf("Name 	=> %s  Types => %T\n", name, name)
	fmt.Printf("Surname => %s  Types => %T\n", surname, surname)
	fmt.Printf("Age 	=> %v  Types => %T\n", age, age)


	const TEMP int= 43    //constant
	//TEMP = 45  invalid

    
}
