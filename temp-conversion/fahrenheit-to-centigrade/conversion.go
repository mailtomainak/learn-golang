/*Write a program that converts
Fahrenheit into Celsius. (C = (F - 32) * 5/9)
from https://www.golang-book.com/books/intro/4*/
package main

import "fmt"

func main() {
	fmt.Println("Enter temperature in Fahrenheit")
	var input float64
	fmt.Scanf("%f", &input)
	output := ((input - 32) * 5 / 9)
	fmt.Println(output)

}
