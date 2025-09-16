package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expense float64
	var tax_rate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expense)

	fmt.Print("Tax rate: ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expense
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("Ratio: ", ratio)

}
