package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	fmt.Println("Go Sorting")

// 	// unsorted array
// 	unsorted := []int{1, 3, 2, 6, 3, 4}

// 	fmt.Println(unsorted)

// 	sort.Ints(unsorted)
// 	fmt.Println(unsorted)
// }

type Programmer struct {
	Age int
}

type byAge []Programmer

func main() {
	programmers := []Programmer{
		Programmer{Age: 30},
		Programmer{Age: 20},
		Programmer{Age: 50},
		Programmer{Age: 1000},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)
}
