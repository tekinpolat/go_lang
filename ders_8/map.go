package main

import (
	"fmt"
	"strings"
)

func init(){
	fmt.Println("Map Tutorials")
}

func main(){
	//Bad
	var counties = []string{"Türkiye", "İtaly", "İspanya"}
	var capitales = []string{"Ankara", "Roma", "Barcelona"}

	counties  = append(counties, "ABD")
	capitales = append(capitales, "Newyork")


	//Good
	countiesCapitales := map[string]string{
		"Türkiye" 	: "Ankara",
		"İtaly" 	: "Roma",
		"İspanya" 	: "Barcelona",
	}

	countiesCapitales["ABD"] = "Newyork"
	countiesCapitales["Azerbeycan"] = "Bakü"

	for country, capital := range countiesCapitales{
		fmt.Printf("Country %s Capital %s\n", country, capital)
	}

	fmt.Println(countiesCapitales["Türkiye"])
	fmt.Println(countiesCapitales["Yunanistan"])

	val, ok := countiesCapitales["Yunanistan"]

	if ok{
		fmt.Printf("%s\n", val)
	}else{
		fmt.Println("Not found")
	}

	val, ok = countiesCapitales["Türkiye"]

	if ok{
		fmt.Printf("%s\n", val)
	}else{
		fmt.Println("Not found")
	}
	
	numbers4 := make(map[int]int) 
	numbers4[0] = 10
	numbers4[1] = 20
	numbers4[2] = 40
	fmt.Printf("%v\n", numbers4)


	info := make(map[string]interface{})
	info["name"] = "Enes"
	info["university"] = "Boğaziçi"
	info["age"]  = 21

	fmt.Printf("%v\n", info)

	for key, val := range info{
		fmt.Printf("%s --> %v\n", key, val)
	}

	delete(info,"age")
	fmt.Printf("%v\n", info)

	/*
	info2 := make(map[interface{}]string)
	info2["Enes"] 		= "name"
	info2["Boğaziçi"] 	= "university"
	info2[21] 			= "age"
	info2[false] 		= "married"

	*/

	info2 := map[interface{}]string{
		"Enes" 		: "name",
		"Boğaziçi" 	: "university",
		21  		: "age",
		false 		: "married",
	}

	info2["Karacaer"] = "surname"

	fmt.Println(strings.Repeat("*", 21))
	for key, val := range info2{
		fmt.Printf("%v <-- %s\n", key, val)
	}

	
	
	

	/* 
		array 
		string 
		numeric  int float 
		interface 
		map 
		struct 
		boolean
	*/
}