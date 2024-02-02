package control

import "fmt"

var (
	x = 0
	y = 0
)

func MoveForward(distance, speed int) {
	fmt.Printf("Adresse mémoire dans MoveForward : %v\n", &distance)
	fmt.Printf("Valeur avant modification : %v\n", distance)

	distance = 10

	fmt.Printf("Valeur après modification : %v\n", distance)

	fmt.Println("Move forward", distance, speed)

	x, y = calculateNextPosition(x, y, distance, 0)
}

func Position() (int, int) {
	return x, y
}

func calculateNextPosition(x, y, distanceX, distanceY int) (newX, newY int) {
	newX = x + distanceX
	newY = y + distanceY

	return
}
