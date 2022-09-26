package service

import "math/rand"

func randomNumberWater() int {
	num := rand.Intn(100-1) + 1
	return num
}

func randomNumberWind() int {
	num := rand.Intn(100-1) + 1
	return num
}
