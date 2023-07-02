//Every Go program it's Included in package

package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go world!")
	distance := 60.8 //distance in km
	fmt.Printf("The distance in miles is %f \n", distance*0.621)

	// Output :-

	//go run {file}.go
	//Hello Go world!
	//The distance in miles is 37.756800

}
