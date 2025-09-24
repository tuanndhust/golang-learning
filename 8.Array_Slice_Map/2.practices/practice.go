package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"football", "movie", "learning"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// sliceH := hobbies[0:2]
	sliceH := hobbies[:2]
	changeH := sliceH[1:3]
	// changeH[0] = hobbies[1]
	// changeH[1] = hobbies[2]
	fmt.Println(changeH)
	// fmt.Println(hobbies)
	fmt.Println("---------------------------")
	goals := []string{"go", "basic"}
	goals[1] = "advanced"
	goals = append(goals, "ez")
	fmt.Println(goals)
	fmt.Println("---------------------------")
	products := []Product{Product{"product1", "1", 300}, Product{"product2", "2", 200}}
	products = append(products, Product{"product3", "3", 350})
	fmt.Println(products)

}
