package main

import (
	"log"

	"xmen-mutant/cmd/api/implement"
)

func main() {
	if err := implement.Run(); err != nil {
		log.Fatal(err)
	}
}
