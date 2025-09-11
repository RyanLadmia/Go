package main

import "testing"

// Fichier de test

// name, a, b, want, isErr
var tests = []struct {
	name  string
	a     float32
	b     float32
	want  float32
	isErr bool
}{
	{"valid", 10, 2.0, 5.0, false}, // On change les valeur pour tester les fonctions
	{"valid2", 10, 2.0, 5.0, false},
	{"invalid", 10, 0.0, 0.0, true},
	{"invalid2", 10, 0.0, 0.0, true},
}

func TestDivide(t *testing.T) {
	// TestDivide est une fonction de test pour la fonction divide.
	// Elle doit commencer par 'Test' pour être reconnue par `go test`.
	// 't *testing.T' permet de gérer les erreurs et les messages de test.

	for _, tt := range tests {
		// On boucle sur chaque cas de test contenu dans le slice 'tests'.
		// '_' signifie qu'on ignore l'index de l'élément, 'tt' est le cas courant.

		want, err := divide(tt.a, tt.b)
		// On appelle la fonction divide avec les valeurs du cas courant.
		// 'want' reçoit le résultat de la division.
		// 'err' reçoit une erreur si la division échoue (ex: division par zéro).

		if (err != nil) != tt.isErr {
			// Vérifie si la présence ou l'absence d'erreur correspond à ce qui est attendu.
			// (err != nil) → vrai si une erreur est survenue.
			// tt.isErr → valeur booléenne attendue du cas de test.

			t.Errorf("%s: got error %v, want error %v", tt.name, err, tt.isErr)
			// Affiche un message d'erreur détaillant :
			// - le nom du test
			// - l'erreur réelle reçue
			// - si une erreur était attendue
		}

		if want != tt.want {
			// Vérifie si le résultat réel correspond au résultat attendu.

			t.Errorf("%s: got %v, want %v", tt.name, want, tt.want)
			// Affiche un message si le résultat est incorrect
			// - le nom du test
			// - la valeur obtenue
			// - la valeur attendue
		}
	}
}

// Lancer les test = go test
