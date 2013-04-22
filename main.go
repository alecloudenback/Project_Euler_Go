package main

import (
	"fmt"
	"os"
)

func main() {
	funcs := map[string]func() string{
		"001": Problem001,
		"002": Problem002,
	}

	fmt.Println("Problem ", os.Args[1], ":\n")
	f, found := funcs[os.Args[1]]
	if found {
		fmt.Println(f())
	} else {
		// do default/error/whatever
		fmt.Println("This problem doesn't exist or hasn't been completed yet.")
	}
}
