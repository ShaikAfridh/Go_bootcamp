package main

import (
	"fmt"
	"math"
)

//we are taking integer 'n' and return the value as boolean

func isPrime(n int) bool {
	//if  number is less than or equal to 1 then its not prime number { returns false}
	if n <= 1 {
		return false
	}
	//loop goes through i=2 ------> sqrt(600851475143) and increment the i value for each loop
	//we run sqrt because product of two factors it will be smaller or equal to the square root so we have taken sqrt of the number
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {

		//if number[n] is divisible by i means it will called as non prime number {returns false}
		if n%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	//we have declared the given number
	number := 600851475143
	//we have declared the largestPrimefactor as equal to 1  beacuse we need to update these value if there is no prime number means it will show as 1
	largestPrimefactor := 1

	// loop goes through i = 2 -------> sqrt(number) and i value increment for each loop
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		//if number is completely divisible by i  and also is it's prime number means we will update the largestPrimefactor with i value
		//also for the loop termination we will divide the number with i for each loop iteation
		if number%i == 0 && isPrime(i) {
			largestPrimefactor = i
			number = number / i

		}
	}

	//incase number is larger than Largest prime factor we will update the value with number in largestPrimefactor
	if number > largestPrimefactor {
		largestPrimefactor = number

	}

	fmt.Printf("The largest Prime factor is %v", largestPrimefactor)
}
