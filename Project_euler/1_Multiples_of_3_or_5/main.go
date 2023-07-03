package main

import "fmt"

func main() {
	//limit of sum of multiples of 3 or 5 is 1000
	limit := 1000
	//we have intiated sum as 0 so we can sum the multiples of 3 or 5 through for loop
	sum := 0
	// loop goes through  i=1 ------> i=1000 and it will increment each time

	for i := 1; i < limit; i++ {
		// we are checking each time value of i it's dividing by 3 or 5
		if i%3 == 0 || i%5 == 0 {
			// we have added the value to sum variable multiples of 3 or 5
			sum = sum + i
		}
	}
	//we have printed the output
	fmt.Printf("Multiples of 3 or 5 upto is %v", sum)
}
