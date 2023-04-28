package main

import "fmt"

func main(){

	name:= "Nawaz"
	fmt.Println(name)

	// variable declaration
	var age int = 20
	fmt.Println(age)

	// taking input
	// var class string
	// fmt.Scanf("%s",&class)

	// arrays
	var fruits [2]int = [2]int{1, 2}
	for index, element := range fruits {
		fmt.Println(index, "->", element)
	}
	fmt.Println(fruits)

	// maps
	code := map[string]string{"Java": ".java", "JavaScript": ".js"}
	fmt.Println(code)
}