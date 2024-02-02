package main

import (
	"fmt"

	"github.com/edafonseca/go-training/functions/control"
)

func main() {
	d := 5

	fmt.Printf("Adresse mémoire : %v, type de base : %T, type de l'opérateur d'adresse : %T \n", &d, d, &d)

	control.MoveForward(d, 10)
	fmt.Printf("Valeur dans la fonction main : %v\n", d)

	x, y := control.Position()

	fmt.Printf("Nouvelle position du robot, x: %v, y: %v\n", x, y)
	control.MoveForward(10, 1)

	x, y = control.Position()
	fmt.Printf("Nouvelle position du robot, x: %v, y: %v\n", x, y)
}
