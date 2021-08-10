package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Al", "Kim", "Belize"}
	messages, err := greetings.HelloMultiple(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
