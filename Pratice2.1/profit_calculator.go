package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expense float64
	// var tax_rate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expense)

	// fmt.Print("Tax rate: ")
	// fmt.Scan(&tax_rate)
	revenue, err1 := getInput("Revenue: ")
	if err1 != nil {
		fmt.Println("Revenue error")
		fmt.Println(err1)
		return
	}
	expense, err2 := getInput("Expense: ")
	if err2 != nil {
		fmt.Println("Expense error")
		fmt.Println(err2)
		return
	}
	tax_rate, err3 := getInput("Tax_rate: ")
	if err3 != nil {
		fmt.Println("Tax_rate error")
		fmt.Println(err3)
		return
	}
	// ebt := revenue - expense
	// profit := ebt * (1 - tax_rate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculatorProfit(revenue, expense, tax_rate)
	writeValueToFile(ebt, profit, ratio)
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("Ratio: ", ratio)

}

func getInput(text string) (float64, error) {
	var output float64
	fmt.Print(text)
	fmt.Scan(&output)
	if output < 0 {
		return -1, errors.New("Value can not be negative numbers")
	}
	if output == 0 {
		return 0, errors.New("Value can not be 0")
	}
	return output, nil
}
func writeValueToFile(ebt, profit, ratio float64) {
	ebtText := fmt.Sprint(ebt)
	profitText := fmt.Sprint(profit)
	ratioText := fmt.Sprint(ratio)
	data := ebtText + "\n" + profitText + "\n" + ratioText + "\n"
	err := os.WriteFile("output.txt", []byte(data), 0644)
	fmt.Println(err)
}
func calculatorProfit(revenue, expense, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
