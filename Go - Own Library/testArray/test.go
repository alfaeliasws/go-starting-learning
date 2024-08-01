package testArray

import (
	"fmt"
	"go-own-lib/alfaArrayProc"
)

func TestSome() {
	// Example usage with a single-dimensional slice
	singleDimensionalSlice := []int{1, 2, 3, 4, 5}
	isEvenSingle := alfaArrayProc.Some(singleDimensionalSlice, func(elem any, i int) bool {
		return elem.(int)%2 == 0 // Check if at least one element is even
	})
	fmt.Println("Some Single-Dimensional Slice (Is at least one element even?):", isEvenSingle)

	// Example usage with a multi-dimensional slice of structs
	type Person struct {
		Name string
		Age  int
	}

	multiDimensionalSlice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "David", Age: 35},
	}

	isOlderThan30 := alfaArrayProc.Some(multiDimensionalSlice, func(elem any, i int) bool {
		return elem.(Person).Age > 30 // Check if at least one person is older than 30
	})
	fmt.Println("Some Multi-Dimensional Slice (Is at least one person older than 30?):", isOlderThan30)

	// Example usage with a map
	sampleMap := map[string]int{
		"one":   1,
		"three": 3,
		"five":  -2,
		"two":   2,
		"four":  -1,
	}

	isGreaterThan3 := alfaArrayProc.Some(sampleMap, func(elem any, i int) bool {
		return elem.(int) > 3 // Check if at least one value is greater than 3
	})
	fmt.Println("Some Map (Is at least one value greater than 3?):", isGreaterThan3)
}

func TestReduce() {
	// Example usage with a single-dimensional slice
	singleDimensionalSlice := []int{1, 2, 3, 4, 5}
	sumSingle := alfaArrayProc.Reduce(singleDimensionalSlice, func(accumulator, currentElement any, index int) any {
		return accumulator.(int) + currentElement.(int) // Sum all elements
	}, 0)
	fmt.Println("Reduced Single-Dimensional Slice:", sumSingle)

	// Example usage with a multi-dimensional slice of structs
	type Person struct {
		Name string
		Age  int
	}

	multiDimensionalSlice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "David", Age: 35},
	}

	averageAge := alfaArrayProc.Reduce(multiDimensionalSlice, func(accumulator, currentElement any, index int) any {
		person := currentElement.(Person)
		averageAge := ((accumulator.(float64) * float64(index)) + float64(person.Age)) / (float64(index) + 1)
		fmt.Println("age:", person.Age)
		fmt.Println(averageAge)
		return averageAge // Accumulate ages
	}, 0.00)

	fmt.Println("Reduced Multi-Dimensional Slice (Average Age):", averageAge)
}

func TestMap() {
	// Example usage with a single-dimensional slice
	singleDimensionalSlice := []int{1, 2, 3, 4, 5}
	mappedSingle := alfaArrayProc.Map(singleDimensionalSlice, func(elem any, index int) any {
		return elem.(int) * 2 // Multiply each element by 2
	})
	fmt.Println("Mapped Single-Dimensional Slice:", mappedSingle)

	// Example usage with a multi-dimensional slice of structs
	type Person struct {
		Name string
		Age  int
	}

	multiDimensionalSlice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "David", Age: 35},
	}

	mappedMulti := alfaArrayProc.Map(multiDimensionalSlice, func(elem any, index int) any {
		person := elem.(Person)
		person.Age++

		return person // Increment the age of each person by 1
	})
	fmt.Println("Mapped Multi-Dimensional Slice:", mappedMulti)
}

func TestSort() {
	// Example usage with a single-dimensional slice
	singleDimensionalSlice := []int{5, 2, 4, 1, 3}
	alfaArrayProc.Sort(singleDimensionalSlice, func(a, b any) bool {
		return a.(int) < b.(int) // Sort in ascending order
	})
	fmt.Println("Sorted Single-Dimensional Slice:", singleDimensionalSlice)

	// Example usage with a multi-dimensional slice of structs
	type Person struct {
		Name string
		Age  int
	}

	multiDimensionalSlice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "David", Age: 35},
	}

	alfaArrayProc.Sort(multiDimensionalSlice, func(a, b any) bool {
		return a.(Person).Age < b.(Person).Age // Sort people by age in ascending order
	})
	fmt.Println("Sorted Multi-Dimensional Slice:", multiDimensionalSlice)

	// Example usage with a map
	sampleMap := map[string]int{
		"one":   1,
		"three": 3,
		"five":  5,
		"two":   2,
		"four":  4,
	}

	alfaArrayProc.Sort(sampleMap, func(a, b any) bool {
		return a.(string) < b.(string) // Sort map keys in ascending order
	})
	fmt.Println("Sorted Map Keys:", sampleMap)
}

func TestFilter() {
	// Example usage with a single-dimensional slice
	singleDimensionalSlice := []int{1, 2, 3, 4, 5}
	filteredSingle := alfaArrayProc.Filter(singleDimensionalSlice, func(elem any, i int) bool {
		return elem.(int)%2 == 0 // Keep only even numbers
	})
	fmt.Println("Filtered Single-Dimensional Slice:", filteredSingle)

	singleDimensionalSliceString := []string{"satu", "dua", "tiga"}
	filteredSingleString := alfaArrayProc.Filter(singleDimensionalSliceString, func(elem any, i int) bool {
		dataString := elem.(string) // Keep only even numbers
		return string(dataString[0]) == "s"
	})
	fmt.Println("Filtered Single-Dimensional Slice:", filteredSingleString)

	// Example usage with a multi-dimensional slice of structs
	type Person struct {
		Name string
		Age  int
	}

	multiDimensionalSlice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "David", Age: 35},
	}

	filteredMulti := alfaArrayProc.Filter(multiDimensionalSlice, func(elem any, i int) bool {
		return elem.(Person).Age > 25 // Keep only people older than 25
	})
	fmt.Println("Filtered Multi-Dimensional Slice:", filteredMulti)

	// Example usage with a map
	sampleMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	filteredMap := alfaArrayProc.Filter(sampleMap, func(elem any, i int) bool {
		return elem.(int) > 1 // Keep only values greater than 1
	})
	fmt.Println("Filtered Map:", filteredMap)
}
