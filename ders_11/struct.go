package main

import "fmt"

type time struct{
	hour,minute,second int
}

type user struct{
	id int
	name,email,phone string
}

func init() {
	fmt.Println("Struct Starting..")
}

func main(){
	// name, surname, age, school 

	pizza := struct{ name string}{ name:"Sucuklu"}

	users := []user{
		{id:1, name:"ali", email: "ali@sirket.com", phone:"5554241212"},
		{id:2, name:"ayşe", phone:"5554241211"},
		{id:3, name:"gamze", email: "gamze@sirket.com", phone:"5554241213"},
	}

	for _, user := range users{
		fmt.Printf("ID:%d Name:%s Email:%s Phone:%s\n", user.id,user.name, user.email, user.phone)
	}

	times := []time{ 
		{hour:10, minute:49, second: 56},
		{hour:24, minute:35, second: 47},
		{hour:12, minute:32, second: 21},
	}

	for _, time := range times{
		fmt.Printf("%d:%d:%d\n", time.hour, time.minute, time.second)
	}


	info := make(map[string]interface{})
	info["name"] 	= "Tekin"
	info["surname"] = "POLAT"
	info["age"]  	= 17

	fmt.Println(info["age"])


	employees := make(map[int]interface{},5)
	employees[0] = map[string]string{
		"name" 		: "Tekin",
		"surname"	: "POLAT",
	}

	employees[1] = map[string]string{
		"name" 		: "Enes",
		"surname"	: "KARACAER",
	}

	for _, employee := range employees{
		fmt.Println(employee)
	}

	type person struct{
		name string
		surname string
		age  int
		ismarried bool
		hobi []string
		schools map[string]string
	}

	//var a  int = 21

	var data1 person = person{
		"Enes", 
		"KARACAER", 
		27, 
		false, 
		[]string{"Kod yazmak", "Spor", "Golf Oynamak"},
		map[string]string{"uni" : "dpu", "lise"	: "Test"},
	}

	fmt.Printf("Ad : %s\n", data1.name)
	fmt.Printf("Hobi 1%s", data1.hobi[0])
	fmt.Printf("Uni %s\n", data1.schools["uni"])

	data2 := person{
		name:"Tekin",
		surname: "POLAT",
		age : 27,
		ismarried: false,
		hobi : []string{},
		schools: map[string]string{"uni":"dpu"},
	}

	fmt.Printf("%#v\n", data2)
}