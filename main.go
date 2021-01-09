package main // files that share same package name can access each others functions

import (
	"fmt"
)

func main() {

	fmt.Println("")                      // Go SDK method names are capitalized
	fmt.Println("Welcome", "Datou", "!") // Print accepts multiple variables and prints with whitespace
	fmt.Println("")

	fmt.Println("=== VARIABLES ===")
	var var1 = "Variables can be declared as - var x = \"y\""
	var2 := "Variables can also be declared as - x := \"y\""
	var3 := "Variables can also explicitly declare type - x string := \"y\""

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println("")

	fmt.Println("=== STRINGS ===")
	str1 := fmt.Sprintf("%s", "String interpolation in Go - fmt.Sprintf(\"%s\", \"\")")

	fmt.Println(str1)
	fmt.Println("")

}
