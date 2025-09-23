package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"3060", "5060", "5090"}
	// productNames = [4]string{"3060", "5060", "5090"}

	prices := [4]float64{37, 38, 38.5, 36.5}

	productNames[0] = "3090"
	// fmt.Println(prices)
	// fmt.Println(productNames)

	fmt.Println(prices[2])
	//slices
	featurePrices := prices[1:3]
	fmt.Println(featurePrices)
	fmt.Println(prices[:2])
	fmt.Println(prices[2:])
	//slices chi la cua so truot nhin vao mot phan mang ban dau
	//thay doi gia tri cua slice cung se thay doi gia tri mang ban dau
	series50 := productNames[1:3]
	series50[0] = "5070"
	fmt.Println(productNames)
	fmt.Println(len(series50), cap(series50))
}
