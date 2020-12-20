package main

import fmt "fmt" 

const (
	One   = 1
	Two   = 2
	Three = 3
)

func main() {
		fmt.Printf("Hello, world;or こんにちは 世界\n")

		var name string = "Bob"
		fmt.Println(name)

		var num int = 1
		fmt.Println(num+num)

		var array1 [4]int
		fmt.Println(array1)

		fmt.Println(One)

		a := 3
		if a ==3 {
			fmt.Println("a = 3")
		} else {
			fmt.Println("a ins not 3")
		}
}