package functions

import (
	"math/rand" // modifier le crypto en math
)

func GenerateRandomNumber(n int) int {
	randomNumber := rand.Intn(n) // Version moderne pour générer un nombre aléatoire.
	//fmt.Println(randomNumber) // Vérification dans le terminal
	return randomNumber
}
