//Go playgrounds:-
//https://go.dev/play/
//https://goplay.space/
package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello World")

	//Declaring a variable

	x := 5.2

	//Go format { It uses gofmt package to format the code }

	//Println function is imported from the fmt package

	//Problem:- calculate (x)^8
  //we need math function to solve the above Problem

	result := math.Pow(x, 8) 

	fmt.Println("x  to the power of 8 is ", result)

	//Go standard library
  // Go provides Huge amount of packages 
  //Below are few pacakges:
  //fmt : this package provides functions for formatting and printing data ,including standard input/output operations.{Println,Printf,Scanf}
  //html/template: The html/template package provides a powerful templating engine for generating HTML content. It offers facilities for creating dynamic HTML templates, rendering data into templates, and preventing common web vulnerabilities.
  //os: The os package provides a platform-independent interface to operating system functionality. It enables tasks such as file system operations, environment variables, process management, and working with command-line arguments.
  // For more about go lang std packages you can visit these website https://pkg.go.dev/std
  
}
