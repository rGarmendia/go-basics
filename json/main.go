package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// Open our JsonFile
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	// var users Users

	// OR WE USE UNSTRUCTURED DATA
	var result map[string]interface{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	// json.Unmarshal(byteValue, &users)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'result' which we defined above
	json.Unmarshal(byteValue, &result)

	// juan := result["users"].(map[string]interface{})

	for _, element := range result["users"].([]interface{}) {
		for k, e := range element.(map[string]interface{}) {
			fmt.Println("key[%s] value[%s]\n", k, e)
		}
	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	// for i := 0; i < len(users.Users); i++ {
	// 	fmt.Println("User Type: " + users.Users[i].Type)
	// 	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
	// 	fmt.Println("User Name: " + users.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }

}
