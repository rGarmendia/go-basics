package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Username  string
	followers int
	Grades    map[string]string
}

type Student struct {
	Firstname, LastName, Email string
	Age                        int
	HeightInMeters             float64
	IsMale                     bool
	Languages                  []string
	Profile                    Profile
}

// A map with key of type strings that receives any type
// type Student map[string]interface{}

func main() {
	// Works with struct, double quotes are not present

	ricardo := Student{
		Firstname:      "Ricardo",
		LastName:       "Garmendia",
		Age:            29,
		HeightInMeters: 1.77,
		IsMale:         true,
		Languages:      []string{"Spanish", "English"},
		Profile: Profile{
			Username:  "ricardog91",
			followers: 1991,
			Grades:    map[string]string{"Math": "A", "Science": "A"},
		},
	}

	//Works with MAPS, look at the double quotes on the keys

	// ricardo := Student{
	// 	"Firstname":      "Ricardo",
	// 	"LastName":       "Garmendia",
	// 	"Age":            29,
	// 	"HeightInMeters": 1.77,
	// 	"IsMale":         true,
	// }

	// encode `ricardo` as JSON
	// ricardoJSON, _ := json.Marshal(ricardo)
	ricardoJSON, _ := json.MarshalIndent(ricardo, "", "  ")

	fmt.Println(string(ricardoJSON))
}
