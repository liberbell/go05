package main

import "fmt"

func main() {
	loons := []string{"bugs", "duffy", "taz"}
	fmt.Printf("loons= %v (type %T)\n", loons, loons)

	fmt.Println(len(loons))

	fmt.Println("------")
	fmt.Println(loons[1])

	fmt.Println("------")
	fmt.Println(loons[1:])
}
