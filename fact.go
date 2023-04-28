package main

import "fmt"

func fact(n int)int{
	if(n==1){
		return 1;
	}

	return n*fact(n-1)
}

func main(){
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &n)
	fmt.Println("The factorial is", fact(n))
}