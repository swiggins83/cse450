package main

import (
	"fmt"
	"os"
	"strconv"
)

var fibValues = map[int]int{
	0: 0,
	1: 1,
}

func fib(i int) int {

	// idea:
	// start the sequence with the base cases predfined in the map (see line 9)
	// check for the values the formula is looking for [fib(i-1), fib(i-2)] in map
	// if they exist, use them for the computation
	// else start an asynchronous function to compute the fib value needed
	// wait for threads (if necessary)
	// return sum of two values discovered

	// value we'll return
	f := 0

	c := make(chan int)
	// to keep track of channels to wait for
	k := 0

	// check for i-1 in map
	if val, ok := fibValues[i-1]; ok {
		f += val
	} else {
		// for channel length
		k++
		// run asynchronous anonymous function to generate fib number
		go func(i int) {
			fibValues[i-1] = fib(i - 1)
			// return new value to channel
			c <- fibValues[i-1]
		}(i)
	}

	// check for i-2 in map
	if val, ok := fibValues[i-2]; ok {
		f += val
	} else {
		// for channel length
		k++
		// run asynchronous anonymous function to generate fib number
		go func(i int) {
			fibValues[i-2] = fib(i - 2)
			// return new value to channel
			c <- fibValues[i-2]
		}(i)
	}

	// wait for return values from channel
	for i := 0; i < k; i++ {
		// add values returned
		f += <-c
	}

	return f
}

func main() {

	n, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fib(int(n)))

}
