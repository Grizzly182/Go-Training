package main

import (
	"fmt"
	"math"
)

func main() {
	const multiplier = 2
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)

	IMT := userWeight / math.Pow(userHeight, multiplier)

	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(IMT)
}
