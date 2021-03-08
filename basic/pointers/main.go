package main

import "fmt"

func yearsUntilRetirment(age int) {
	fmt.Println(100 - age)
}
func main() {
	var age *int
	age = new(int)
	*age = 26

	yearsUntilRetirment(*age)
}

// // Examples

// func testFunc(id string) (Guitarist, error) {
//     result, err := getSongs(id)
//     if err != nil {
//         return Guitarist{}, err
//     }

//     return result, nil
// }

// // vs Pointers

// func testFunc(id string) (*Guitarist, error) {
//     result, err := getSongs(id)
//     if err != nil {
//         return nil, err
//     }

//     return result, nil
// }
