package main

import "fmt"

func main() {
	//gofmt ----> go formatter it will format the source code it will help us to read easily and understad
	//gofmt is built in ruuntime, it format go code into set of stable and well understood language rules
	// we can also run manually at the command line
	//EX:- gofmt -w main.go
	// Before:-
	// 		fmt.Println("Hello world!")
	// distance := 60.8
	// fmt.Printf("The distance in miles is %f \n", distance*0.62137)

	fmt.Println("Hello world!")
	distance := 60.8
	fmt.Printf("The distance in miles is %f \n", distance*0.62137)

}
