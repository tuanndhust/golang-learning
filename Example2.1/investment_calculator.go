package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64 = 1000
	// var investmentAmount, years = 1000, 10
	var expectedReturnRate = 5.5
	// expectedReturnRate := 5.5
	// var years float64 = 10
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculorFutureValue(investmentAmount, expectedReturnRate, years)

	// var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	// fmt.Printf("Future Value: %v\nFuture Real value: %v\n", futureValue, futureRealValue)
	fmt.Printf(`Future Value: %.0f
	Future Real value: %.0f`, futureValue, futureRealValue)

	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Real value: %.2f\n", futureRealValue)
	// fmt.Print(formattedFV, formattedRFV)

}

//function example

func outputText(text string) {
	fmt.Print(text)
}

//	func calculorFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
//		fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
//		rfv := fv / math.Pow(1+inflationRate/100, years)
//		return fv, rfv
//	}
func calculorFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}
