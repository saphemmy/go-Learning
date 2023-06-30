package main

import "fmt"

var sayhi string = "hello world"

//Main / Parent function
func main()  {
	testing()
	testing2()
}

//child function
func testing()  {
	fmt.Println(sayhi)
}

//Another child function
func testing2()  {
	fmt.Println(sayhi + " again")
}