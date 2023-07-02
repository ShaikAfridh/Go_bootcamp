
package main

import "fmt"



func main() {
	
	// Output :-

	//go run {file}.go
	//Hello Go world!
	//The distance in miles is 37.756800

	//Compiling (go build) and Running Go Applications (go run)

	//go is tool to manage source code of GO

	//go run : It compiles + run the application { it doesn't create executable }.

	//go run -x file.go :- it will display execution process of program

	//go build : It Only compiles the application and produce executable
	// go build file.go

	// $ go build
	// $ ls
	// main  main.go
	// $ ./main
	// Hello Go world!
	// The distance in miles is 37.756800
	// afridi@ubuntu:~/Desktop/go_bootcamp/day_1/application_structure$

	//If you want to give executable file another name
	//go build -o {Executable_filename}

	// $ go build -o  myapp
	// $ ls
	// main  main.go  myapp
	// $ ./myapp
	// Hello Go world!
	// The distance in miles is 37.756800

	// Compiling program for different OS
	//Windows :- GOOS=windows GOARCH=amd64 go build -o winapp.exe
	//Linux :- GOOS=linux GOARCH=amd64 go build -o linuxapp
	//Mac :- GOOS=darwin GOARCH=amd64 go build -o macapp

	//Windows:-
	// $ GOOS=windows GOARCH=amd64 go build -o winapp.exe
	// $ ls
	// main  main.go  myapp  winapp.exe

	//Go install

	//go install & go build both compile the package in current directory
	//Package is "main" then with go build it will place result of executable in current directory
	//But go install moves the executable to GOPATH/bin
	// while runing go install we need to use relative path to GOPATH/src


// $ go install
// $ls  ~/go/bin/
// application_structure dlv  gomodifytags  goplay  gopls  gotests  impl  staticcheck



}
