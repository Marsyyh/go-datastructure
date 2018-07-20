package main

import "fmt"

func main() {
	t := [3]int{1, 2, 3}
	fmt.Println(len(t))
	fmt.Println(cap(t))
	// t = append(t, 3)
	// fmt.Println(len(t))
	// fmt.Println(cap(t))
	// t = append(t, 3)
	// t = append(t, 3)
	// t = append(t, 3)
	// fmt.Println(len(t))
	// fmt.Println(cap(t))
}
