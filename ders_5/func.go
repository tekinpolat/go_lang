package main

import "fmt"

func init(){
	fmt.Println("Function starting..")
}

func init() {
    fmt.Println("Second init function")
}


func main(){

	//var number int 

	//var number1 = 35
	//var number2 int = 43

	test()
	sum(21, 43)
	sum1(21, 43, 35, "Enes")

	total := sum2(21, 35)
	fmt.Printf("%d + %d = %v\n", 21, 35, total)

	numbers := []int{21, 35, 43}

	var productResult int = product(numbers)
	fmt.Printf("Product => %d\n", productResult)


	var arr = myRange(9)

	fmt.Printf("%v\n", arr)

	value1, value2 := multiReturn(5)
	fmt.Printf("Value 1 => %d\n", value1)
	fmt.Printf("Value 2 => %d\n", value2)

	sum3(21,35)
	sum3(21,35, 43)
	sum3(21,35, 43, 2)
	sum3(21,35, 43, 2, 5)
}

func sum3(numbers ...int) int{
	total := 0 

	fmt.Printf("%T\n", numbers)

	for _, val := range numbers{
		total += val
	}

	return total
}

func multiReturn(number int) (int, int){
	return number*5, number*10
}

func myRange(number int) []int{
	arr := make([]int, number, number)    //9  0 - len   number - capacity

	fmt.Printf("%v\n", arr)
	fmt.Printf("Length %d\n", len(arr))
	fmt.Printf("Capacity %d\n", cap(arr))

	for index := 1; index <= number; index++{
		arr[index-1] = index*10
	}

	return arr
}

func product(numbers []int) int{
	var productResult int = 1 

	for _, val := range numbers{
		productResult *= val
	}

	return productResult
}

func test(){
	fmt.Println("test function non parameters non return")
}

func sum(number1 int, 	number2 int){
	total := number1 + number2

	fmt.Printf("%d\n", total)
}

func sum1(number1, number2, number3 int, name string){
	var total int = number1 + number2 + number3

	fmt.Printf("%d\n", total)
}

func sum2(number1, number2 int) int{
	return number1 + number2
}

