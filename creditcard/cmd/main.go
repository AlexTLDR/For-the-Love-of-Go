package main

import (
	"creditcard"
	"fmt"
	"log"
)

func main() {
	card, err := creditcard.New("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(card.Number())
}
