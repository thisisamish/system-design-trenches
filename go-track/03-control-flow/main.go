package main

import "fmt"

func main() {
	fmt.Println("Grade(85):", Grade(85))
	fmt.Println("SumTo(10):", SumTo(10))
	fmt.Println("Reverse(\"hello\"):", Reverse("hello"))

	fmt.Print("Primes below 30:")
	for n := 0; n < 30; n++ {
		if IsPrime(n) {
			fmt.Print(" ", n)
		}
	}
	fmt.Println()
}
