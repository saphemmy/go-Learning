package main

import "fmt"

func main()  {
	devops_topics := [2]string{"Go", "VS Code"}

	fmt.Println(devops_topics)

	// Create a slice that only show the second value
	// the_slice := devops_topics[1:2]

	// fmt.Println(the_slice)

	new_slice := devops_topics[0:2]
	fmt.Println(new_slice[0])
	fmt.Println(new_slice[1])

	fmt.Println("we are studying " + new_slice[0] + " in " + new_slice[1])
}