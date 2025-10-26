package main

import (
	"fmt"
	"math"
)

func main() {
	IMT := calculateIMT(getUserInput())
	outputIMTResult(IMT)
}

func calculateIMT(userHeight float64, userWeight float64) float64 {
	const multiplier = 2
	return userWeight / math.Pow(userHeight/100, multiplier)
}

func outputIMTResult(imt float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", imt)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
