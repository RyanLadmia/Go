package main

import (
	"fmt"
	"project-go/lesson19/functions"
)

/*

Les channels :
Pour envoyer des infos d'une partie de notre programme à une autre partie.
On a un nombre limité de channel

*/

func CalculateValue(intChan chan int) {
	randomNumber := functions.GenerateRandomNumber(50)
	intChan <- randomNumber
}

func main() {
	//functions.GenerateRandomNumber(100) // On genere un nombre aléatoire entre 0 et 99
	foo := make(chan int)
	defer close(foo) // execute le code apres defer lorsque le code main est terminé, utile pour fermer l'accès à la base de donnée

	go CalculateValue(foo) // On utilisa une routine pour appeler un channel avec le mot clé go

	num := <-foo // On stoxke les infos du channel dans num

	fmt.Println(num) // On print le nombre aléatoire
}
