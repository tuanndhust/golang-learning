package main

import "fmt"

func main() {
	websites := map[string]string{ //[type of keys] type of values
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	//add key-value
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	//remove item
	delete(websites, "Google")
	fmt.Println(websites)
}
