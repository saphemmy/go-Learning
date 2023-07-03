package main

import (
	"fmt"
	"log"
)

func main()  {
	ifStatement()
}Mike

func ifStatement()  {
	test := "Mike"

	if test != "Mike" {
		log.Fatalln("Its broken!")
	}

	if test == "Mike" {
		fmt.Println("Hurray!")
	}
}