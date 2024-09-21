package main

import "fmt"

func main() {
	fmt.Print("Hello World!\n")
	var i int = 9
	fmt.Printf("%d is a number\n", i)
	var f float32 = 3.14
	fmt.Println(f)
	var b bool = true
	fmt.Println(b)
	var s string = "Don't worry about apostrophes"
	fmt.Println(s)
	var favNums [3]int
	favNums[0] = 1
	favNums[1] = 2
	favNums[2] = 3
	fmt.Println(favNums)
	Nums := [4]int{50, 25, 30, 33}
	fmt.Println(Nums)
	var exampleSlice []int
	exampleSlice = append(exampleSlice, 1)
	exampleSlice = append(exampleSlice, 2)
	exampleSlice = append(exampleSlice, 3)
	fmt.Println(exampleSlice)
	fmt.Printf("%d is the length of the Slice\n", len(exampleSlice)) // prints 3
	fmt.Printf("%d is the capacity of the underlying array\n", cap(exampleSlice))
	makeSlice := make([]int, 0, 5)
	fmt.Println(makeSlice)

	if userAge := 26; userAge < 20 {
		fmt.Println("Below 20")
	} else if userAge >= 20 && userAge <= 60 {
		fmt.Println("Between 20 and 60")
	} else {
		fmt.Println("Above 60")
	}
}
