package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int
	fmt.Printf("a : %T,  b : %T", a, b)
	testArray := [...]int{1:1, 5:2, 10:3}
	fmt.Println()
	fmt.Println(testArray)
	fmt.Println(len(testArray))

	for i := 0; i < len(testArray); i++ {
		fmt.Print(testArray[i])
	}
	fmt.Println()
	fmt.Println("========================")

	for i,v := range testArray {
		fmt.Println(i, v)
	}
	fmt.Println()
	fmt.Println("========================")

	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _,v := range arr {
		sum+= v
	}
	fmt.Printf("arr sum = %d", sum)

	fmt.Println()
	fmt.Println("========================")

	for i, v := range arr {
		fmt.Println(i, v)
	}

}