package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {

	n, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println(err)
	}

	f := fib()
	for i := 0; i < int(n); i++ {
		fmt.Println(f())
	}

}
