package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Write your code here
	bubblegum := 202
	toffee := 118
	iceCream := 2250
	milkChocolate := 1680
	doughnut := 1075
	pancake := 80
	var income = bubblegum + toffee + iceCream + milkChocolate + doughnut + pancake
	var staff_expenses int = 0
	var other_expenses int = 0
	fmt.Println("Earned amount:\n" +
		fmt.Sprintf("Bubblegum: $%d\n", bubblegum) +
		fmt.Sprintf("Toffee: $%d\n", toffee) +
		fmt.Sprintf("Ice cream: $%d\n", iceCream) +
		fmt.Sprintf("Milk chocolate: $%d\n", milkChocolate) +
		fmt.Sprintf("Doughnut: $%d\n", doughnut) +
		fmt.Sprintf("Pancake: $%d\n", pancake))
	fmt.Println("Income: $" + strconv.FormatFloat(float64(income), 'f', 1, 32))
	fmt.Println("Staff expenses: ")
	fmt.Scan(&staff_expenses)
	fmt.Println("Other expenses: ")
	fmt.Scan(&other_expenses)
	fmt.Printf("Net income: $%d", income-staff_expenses-other_expenses)
}
