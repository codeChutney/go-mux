package main

import (
	"fmt"
)

func main() {
	fmt.Println("Entered main")

	a := Application{}
	//lets just use some garbage values for now
	a.Initialize("postgres", "", "postgres")
	a.Run(":8010")

	fmt.Println("Exiting main")
}
