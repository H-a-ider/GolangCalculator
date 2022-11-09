package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {

	var x float64
	pflag.Float64VarP(&x, "firstnumber", "x", 0.00, "This is the first Number")
	var y float64
	pflag.Float64VarP(&y, "secondnumber", "y", 0.00, "This is the second Number")
	var action string
	pflag.StringVarP(&action, "action", "a", action, "Given Action is Done")
	pflag.Parse()

	switch action {
	case "add":
		fmt.Println("Addition of first and second number is", add(x, y))
	case "sub":
		fmt.Println("Subtraction of first and second number is", sub(x, y))
	case "mul":
		fmt.Println("Multiplication of first and second number is", mul(x, y))
	case "div":
		fmt.Println("Division of first and second number is", div(x, y))

	}

}

// type action struct {
// 	x float64
// 	y float64
// }
