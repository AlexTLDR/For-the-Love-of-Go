package main

import "fmt"

func Overwrite(s []int) {
	s[0] = 0
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	Overwrite(nums)
	fmt.Println(nums)
}
