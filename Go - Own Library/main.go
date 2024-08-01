package main

import (
	"fmt"
	testArray "go-own-lib/testArray"
)

func main() {
	fmt.Println("SOME")
	testArray.TestSome()
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("REDUCE")
	testArray.TestReduce()
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("MAP")
	testArray.TestMap()
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("SORT")
	testArray.TestSort()
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("FILTER")
	testArray.TestFilter()
	fmt.Println("\n")
	fmt.Println("\n")
}
