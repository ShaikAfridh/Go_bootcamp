package main

import "fmt"

func main() {
	//Variable:- it's a name  for memory location where a value of a specific type{int,float,string etc..} is stored

	//In Go variable created at runtime

	//There are two ways of declaring variable

	//1. Using the var keyword

	//int age = 10;  //C++ Variable declaration

	var age int = 30
	fmt.Println("Age: ", age) //Age:  30

	//Go automatically assign the type based on intial value {"Dan"}
	var name = "Dan"
	// fmt.Println("Your name is : ", name) //Your name is :  Dan

	//Note:-Unused variable throws errors while compilation

	_ = name //to ignore these type unused variable error we can use blank identifier{"_"}

	//2. Short Declaration [ := ]
	s := "Learning golang!"
	fmt.Println(s) //Learning golang!
	//It's only work on block scope

	//we can create one or more varibale with ":="

	// name := "Andrei" //we can't use colon equal for already defined variables { Error:- no new variables on left side of :=}

	//In case if i want to use the variable we can use "="
	name = "Andrei"
	fmt.Println(name) //Andrei

	name1 := "Jhon"
	_ = name1 //for these short declaration varibles also we can use blank identifier "_"

	//Multiple Declaration

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// car,cost := "BMW", 60000 //we can't reassign  already two existing  varibales{car,cost} at a time it will throws error

	car, year := "BMW", 2018 //we can assign value with combination of Already exist variable + new variable with ":="

	_ = year

	//Boolean
	var opened = false
	//used variable{bool} + new varibale{string}
	opened, file := true, "a.txt"

	_, _ = opened, file

	//Multiple variable declaration
	//1.
	var (
		salary    float64
		firstname string
		gender    bool
	)

	fmt.Println(salary, firstname, gender) //0  false //Automatically take O value for numeric type in golang

	var d, e, f int
	fmt.Println(d, e, f)

	//2.
	var i, j int
	i, j = 5, 8

	// _,_ = i,j

	j, i = i, j       //we can swap two varible values
	fmt.Println(i, j) //8 5

	sum := 5 + 2.3
	fmt.Println(sum)

	//Types and Zero vallues

	var a = 4   //int
	var b = 5.2 //float
	// a = b //type of varible can't be changed Once declared !{ Error:- cannot use b (variable of type float64) as int value in assignmen}

	// Incase i want to Assign the value of b to a we need convert type {float ----> int} [ type casting]
	a = int(b)

	fmt.Println(a, b)
	fmt.Println()

	// we can't assign the variable to another type {Error:- cannot use "5" (untyped string constant) as int value in assignment}
	// var x int //integer
	// x = "5" //string literal["5"] //Literal:- value itself

	//Note:- In Go all variables are Intialized

	//Go Zero Values:
	//      	 - Numeric types: 0
	//			 - bool type: false
	//			 - string type: ""(empty string)
	//			 - pointer type: nil

	var value int
	var price float64
	var company string
	var profit bool
	fmt.Println(value, price, company, profit) //0{int} 0{float64}  {string} false{bool}

	//Comments:- it's not execute It's good practise to understand the code

	//There are two types of Comments

	//1. this is a singleline comment [//]

	//declaring a new variable of type int
	number := 18
	_ = number

	//2.this is multiline comment [/* -------------code ------------*/]
	/*
		number := 18
		_ = number
	*/







	//Naming Conventions in GO


	//Rules:-

	//Names of variables,functions,Constants and types , statement lables and Packages follow very simple rule Names start with "letter" and "Underscore(_)"

	//In Go there are 25 keywords that can't be used as names
	/*
	break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
*/

	//Case Matters:- quickSort and Quicksort are different variables

	//Use first letters of the words



	var mv int //max value
	var max_value int //NOT idiomatic {It's Ok for Python}

	var packetsReceived int //Not OK, name toolong

	var b int //OK
	
	const MAX_VALUE = 100 //NOT OK
	const N = 100 // Better

	var Max = 100 // Go has special feature which variable start with Captial letter can be exported into other packages
	var max = 100 //local variable
	maxValue := 1000 //recommeneded
	max_Values := 1000 //Not recommeneded

	writeToDB := true // ok, idiomatic
	wrtieToDb := false //not ok





}
