package main

import (
	"fmt"
	"math"
)

func main() {
	const multiplier = 2
	userHeight := 1.77
	userWeight := 65.0
	IMT := userWeight / math.Pow(userHeight, multiplier)

	fmt.Print(IMT)
}
