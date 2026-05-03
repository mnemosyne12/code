package helpers

import "math/rand"

func RandomNumber(n int) int {
	value := rand.Intn(n)

	return value

}

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber
}
