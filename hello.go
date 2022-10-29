package main

import (
	"fmt"
	"log"

	madmod "github.com/madpin/madmod"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := madmod.Hello("random_name_86")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
