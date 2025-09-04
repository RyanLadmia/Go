package main

import "fmt"

/*
	Les array : collections de données
	On les déclare comme ceci :

	var list [2]int (on stocke des entiers)
	list[0] = 10 (le nombre 0 est la première position)
	list[1] = 20

	ou

	newList := [...]int{20, 50} (Ici on déclare et assigne en même temps)

	ou

	bigList := [...]int{40, 50, 69, 420, 7777, 50000}
	for k, v := range bigList {
		fmt.Printf("Position %d est égal à %d. \n", k, v)
	}
*/

func main() {

	// 1°)

	var list [3]int
	list[0] = 10
	list[1] = 20
	list[2] = 30

	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(list[2])

	fmt.Println("--------")

	newList := [2]int{40, 50} // On peut mettre ... à la place de 2 et le code compte auto les éléments
	fmt.Println(newList)
	fmt.Println(newList[0])
	fmt.Println(newList[1])

	fmt.Println("--------")

	bigList := [...]int{10, 40, 50, 69, 420, 777, 500000}
	for pos, value := range bigList { // pos = variable nommée de position (ici 0) et value représente la valeur dans le tableau
		fmt.Printf("Position %d est égale à %d.\n", pos, value) // %d sert à affichée une valeur décimale
	}
}
