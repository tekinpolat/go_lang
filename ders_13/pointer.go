package main

import "fmt"

func init(){
	fmt.Println("Pointer Tutorial")
}

func changeValue(number *int){
	*number = 43
}



func main(){
	var number1 int = 21   // &number1
	fmt.Printf("%p\n", &number1)

	var number1Ptr *int 
	number1Ptr = &number1

	//&number1Ptr  değişkenin kendi adresi
	//*number1Ptr   
	//number1Ptr

	fmt.Printf("%p\n", number1Ptr)

	changeValue(&number1)

	fmt.Printf("%d\n", number1)


	var number2 float32 = 54.3

	var number2Ptr *float32

	number2Ptr = &number2

	fmt.Printf("%.1f\n", *number2Ptr)

	var numbers = []int{1, 2, 3, 4}
	arrPass(numbers)    //numbers   
	
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%p\n", &numbers[0])
	fmt.Printf("%p\n", numbers)

	var name string = "Tekin"
	name[0] = "M"
	fmt.Printf("%s\n",name)

}


func arrPass(numbers []int){
	numbers[0] = 21
}