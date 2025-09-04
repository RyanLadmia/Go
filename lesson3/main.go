package main

import "fmt"

/*
Les types :
GO possède 17 types différents, dont 15 sont numériques.
booéléen : bool (true ou false)
chaine de caractères : string ("texte")
nombres (15 types)
	entiers signés INT (positifs et négatifs) : int8, int16, int32, int64, int (au moins 32 bits)
	entiers non signés UINT (positifs uniquement) : uint8, uint16, uint32, uint64, uint (au moins 32 bits)
	nombres à virgule flottante (décimaux) : float32, float64
	nombre complexe : complex64, complex128
byte : alias pour uint8
rune : alias pour int32, représente un point de code Unicode

Le type int est généralement utilisé pour les entiers, et le type float64 pour les nombres à virgule flottante.
Le type string est utilisé pour les chaînes de caractères.
Le type bool est utilisé pour les valeurs booléennes (true ou false).

Les variables peuvent être déclarées avec le mot clé var, suivi du nom de la variable, du type et éventuellement d'une valeur initiale.
Les variables peuvent aussi être déclarées avec l'opérateur :=, qui infère le type de la variable à partir de la valeur initiale.

!!!!! Attention : en GO, une variable déclarée mais non utilisée provoque une erreur de compilation !!!!!
!!!!! nil est la valeur zéro spéciale qui représente l’absence de valeur. C’est l’équivalent de null en Java / JS ou None en Python.
!!!!! Mais en Go, nil ne peut pas s’appliquer à tous les types : seulement aux pointeurs, aux interfaces, aux fonctions, aux slices, aux maps et aux channels.
!!!!! Pas pour les types de base comme int, float, bool, string, etc.
!!!!! En GO, les variables non initialisées prennent automatiquement la valeur zéro de leur type :

*/

func main() {

	var (
		b   bool
		s   string
		i   int
		u   uint    // unsigned int (positif uniquement)
		u8  uint8   // Le 8 bloque l'assignation aux valeurs entre 0 et 255 (8bits)
		i8  int8    // Le 8 bloque l'assignation aux valeurs entre -128 et 127 (8bits)
		i16 int16   // Le 16 bloque l'assignation aux valeurs entre -32768 et 32767 (16bits)
		u16 uint16  // Le 16 bloque l'assignation aux valeurs entre 0 et 65535 (16bits)
		f   float32 // Le type float32 et float64 sont utilisés pour les nombres à virgule flottante (décimaux)

		// Le type int et uint sont par defaut de 32 ou 64 bits selon l'architecture de la machine (32 ou 64 bits) ici 64 bits pour un MacBook air M4
		// Il n'y a pas de type float par defaut (il faut choisir entre float32 et float64)
		// float32 est plus rapide mais moins précis que float64 (7 chiffres significatifs contre 15 pour float64)
	)

	b = true
	s = "Batman"
	i = -16
	u = 42
	u8 = 255
	i8 = -127
	i16 = 21500
	u16 = 65535
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, i16, u16, f)

}
