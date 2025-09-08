package main

import "fmt"

/*

Découverte des interfaces avec GO :
une interface est un type qui définit un ensemble de méthodes qu’un autre type doit posséder pour “satisfaire” (implémenter) cette interface.
	•   Une interface décrit un comportement (ce que quelque chose peut faire).
	•	Une struct (ou autre type) fournit la mise en œuvre (comment elle le fait).
*/

type Animal interface {
	Noise() string
	NummberOfLegs() int
}

type Cat struct {
	Name  string
	Breed string
}

type Spider struct {
	Name     string
	Breed    string
	Venomous bool
}

func main() {
	cat := Cat{
		Name:  "Miaouw",
		Breed: "Miaous",
	}
	printAnimalInfo(&cat) // & : car il y a un pointer.

	spider := Spider{
		Name:     "Spidey",
		Breed:    "Mimigal",
		Venomous: true,
	}
	printAnimalInfo(&spider)
}

func printAnimalInfo(a Animal) {
	fmt.Println("Le bruit de cet animal est", a.Noise(), "et il possède", a.NummberOfLegs(), "pattes.")
}

func (c *Cat) Noise() string { // Les receiver sont de type pointer
	return "Miaou"
}

func (c *Cat) NummberOfLegs() int {
	return 4
}

func (s *Spider) Noise() string {
	return "None"
}

func (s *Spider) NummberOfLegs() int {
	return 8
}
