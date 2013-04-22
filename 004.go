package main

import (
	"fmt"
	"runtime"
)

func isPalindrome(num int) bool {
	var n, reverse, dig int
	n = num
	reverse = 0
	for n > 0 {
		dig = n % 10
		reverse = reverse*10 + dig
		n = n / 10
	}
	return num == reverse
}

// given num, traverse the integers less than num to find largest palindrome mulitple
func traversePalindrome(num int) int {
	a := num
	b := num
	n := num
	for !isPalindrome(n) {
		n = a * b
		b -= 1
	}
	return n // return the given limit num and palindrome n
}

func palindromesUnder(num int) *[]int {
	palindromes := make([]int, num)

	c := make(chan int)

	for i := 0; i < num; i++ {
		go func(i int) {
			palindromes[i] = traversePalindrome(i + 1)
			c <- i // don't have to give channel number, but more intuitive than nothing
		}(i)
	}

	// wait for goroutines to finish
	for i := 0; i < num; i++ {
		<-c
	}
	return &palindromes
}

func Problem004() string {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.

	p := *palindromesUnder(1000)
	fmt.Println(p)

	//find maximum
	max := p[0]

	for _, xi := range p {
		if xi > max {
			max = xi
		}
	}

	return fmt.Sprint(max)
}
