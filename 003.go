package main

import "fmt"

func isPrime(n int) bool {
	limit := n / 2
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// return the first prime int and the reduced number
func lowestPrimeFactor(num int) int {
	limit := num / 2
	for i := 2; i < limit; i++ {
		if isPrime(i) && num%i == 0 {
			return i
		}
	}
	return num
}

// return a slice of the prime factors
func Problem003() string {
	num := 600851475143
	return fmt.Sprint(primeFactors(num))
}

func primeFactors(num int) []int {
	if isPrime(num) {
		return []int{num}
	} else {
		fact := lowestPrimeFactor(num)
		return append(primeFactors(num/fact), fact)
	}
}
