package main

import "fmt"

func TwoTimes(n int) int {
	return 2 * n
}

func ThreeTimes(n int) int {
	return 3 * n
}

func main() {
	var f func(n int) int;
	f = TwoTimes
	fmt.Printf("%#v on 10 is %d\n", f, f(10))
	f = ThreeTimes
	fmt.Printf("%#v on 5 is %d\n", f, f(5))
}
