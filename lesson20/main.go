package main

import (
	"encoding/json"
	"fmt"
)

/*

Manipulation d'un JSON :

*/

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	jsonFromAPi := `
[
	{
		"name": "Ryan",
		"age": 28,
		"email" : "bat@example.com",
		"phone": "0618273647"
	},
	{
		"name": "Youssef",
		"age": 33,
		"email" : "you@example.com",
		"phone": "0638291836"
	}
]`

	var users []User // on utilise un slice parce qu'on a plusieurs entrées et infos dans le json

	err := json.Unmarshal([]byte(jsonFromAPi), &users)
	if err != nil {
		fmt.Println("error unmarshalling json:", err)
	}
	fmt.Printf("json: %v\n", users)

	// Transformer le slice en JSON -----------------------------------------

	var myStruct []User // On crée la structure slice[]

	var user_one User // Création de l'utilisateur 1
	user_one.Name = "Batman"
	user_one.Age = 29
	user_one.Email = "batman@example.com"
	user_one.Phone = "0671263748"

	myStruct = append(myStruct, user_one) // Ici on rajoute les infos dans le slide myStruct

	var user_two User // Création de l'utilisateur 2
	user_two.Name = "Robin"
	user_two.Age = 15
	user_two.Email = "robin@example.com"
	user_two.Phone = "0671263700"

	myStruct = append(myStruct, user_two)

	jsonFromSlice, err := json.MarshalIndent(myStruct, "", " ") // indent et "" " " pour les indentations dans le terminal.
	if err != nil {
		fmt.Println("error unmarshalling json:", err)
	}
	fmt.Println(string(jsonFromSlice)) // string pour convertir le message en string
}
