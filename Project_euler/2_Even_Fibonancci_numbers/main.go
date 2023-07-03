package main

import "fmt"

func main() {
	//for these problem they have given limit to 4 million
	limit := 4000000
	//first number[num1] is set to 1
	//second number[num2] is set to 2
	//we have intiated the sum variable with 0

	num1, num2, sum := 1, 2, 0
	//for loop goes until secondnumber[num2] greater than or equal to 4 million then loop breaks
	for num2 <= limit {
		//we need only sum of fibonancci numbers
		if num2%2 == 0 {
			//we have added the num2 to the sum each time
			sum += num2
		}
		//we have changed num1 with num2 and num2 with addition of num1 + num2 because we need sum of previous two numbers of sum
		num1, num2 = num2, num1+num2

	}
	fmt.Printf("Even fibonacci numbers upto 4 milion is %v", sum)
}
