package main

import "fmt"

func main() {
	var cost float32
	var tip, tax int
	fmt.Scanf("%f\n", &cost)
	fmt.Scanf("%d", &tip)
	fmt.Scanf("%d", &tax)

	mealCost := cost + (cost * float32(tip) / 100.0) + (cost * float32(tax) / 100.0)

	fmt.Printf("The total meal cost is %0.0f dollars.\n", mealCost)
}
