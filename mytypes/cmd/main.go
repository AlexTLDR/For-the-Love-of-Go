package main

import (
	"fmt"
	"mytypes"
)

func main() {
	var mb mytypes.MyBuilder
	mb.WriteString("hello")
	fmt.Println(mb.String())
}
