package main

import "fmt"

func main()  {
	
	type car struct {
		make 		string
		model		string
		color 		string
		quantity	int
	}

	// You can initialize your Struct in a few ways
	// var iWantAnR8 car
	// iWantAnR8 := new(car)

	// this is the shortHand declaration with it's Keys
	iWantAnR8 := car{
		make: 		"Audi",
		model:		"R8",
		color:		"blue",
		quantity:	1,
	}

	fmt.Println(iWantAnR8)
}