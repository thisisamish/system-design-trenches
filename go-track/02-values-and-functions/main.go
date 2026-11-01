package main

import "fmt"

func main() {
	fmt.Println("100°C in °F:", CelsiusToFahrenheit(100))

	dollars, cents := SplitCents(1999)
	fmt.Println("1999 cents is", dollars, "dollars and", cents, "cents")

	fmt.Println("50 of 200 is", PercentOf(50, 200), "percent")
	fmt.Println("90 seconds is", MillisFromSeconds(90), "ms")
	fmt.Println("average of 10 and 15:", Average(10, 15))
}
