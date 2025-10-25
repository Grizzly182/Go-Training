package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.77
	userWeight := 65.0
	IMT := userWeight / math.Pow(userHeight, 2)

	fmt.Print(IMT)
}
