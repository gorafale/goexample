//Go Program execution starts from main package first line of code
package main

//factored import statement
//import required package
import (
	"fmt"
)

//init function will be excuted implicitly. no need to call it
//it wil be executed once regardless of how many time package is imported
//Unlike other functions, this function can be defined any no of time in go program
//In case of multiple occurance, these will be excuted in they are defined
func init() {
	fmt.Println("Define Before Main Function - 1")
}

func init() {
	fmt.Println("Define Before Main Function - 2")
}

//main function
//Its mandatory to define main function to start program execution (try by renmaing or remove main method defined below)
//only declaration statement (ex variable, function declartion etc) allowed to write outside.
//init function will be executed before main irespective of its defined order
//rest of the program statement will be executed in defined order
func main() {
	fmt.Println("Main function")
}

func init() {
	fmt.Println("Define After Main Function - 1")
}

/*
Run Program
go run main.go
Define Before Main Function - 1
Define Before Main Function - 2
Define After Main Function - 1
Main function

*/
