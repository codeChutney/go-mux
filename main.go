package main

import (
	"fmt"
)

func main() {
	fmt.Println("Entered main")

	a := Application{}
	//lets just use some garbage values for now
	a.Initialize("postgres", "", "postgres")

	fmt.Println("Exiting main")
}
