package main

import "fmt"

func main()  {
	name("Femi")
	double_name("Bunmi", "Bayo")
}

func name(name string)  {
	fmt.Println(name)
}

func double_name(name1, name2 string)  {
	fmt.Println(name1 + " is friends with " + name2)
}