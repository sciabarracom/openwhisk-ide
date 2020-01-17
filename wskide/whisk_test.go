package main

import "fmt"

func ExampleWhiskParse() {
	fmt.Println(1, whiskParse("whisk"))
	fmt.Println(2, whiskParse("whisk deploy"))
	fmt.Println(3, whiskParse("whisk destroy"))
	// Output:
	// 1 false
	// Deploying Whisk
	// 2 true
	// Destroying Whisk
	// 3 true
}
