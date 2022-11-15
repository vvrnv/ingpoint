package main

import (
	"fmt"
	"log"

	"github.com/bitfield/script"
)

func main() {
	query := fmt.Sprintln(".items[].spec.rules[].host")
	_, err := script.Exec("kubectl get -A ing -ojson").JQ(query).Freq().Column(2).Stdout()
	if err != nil {
		log.Fatal(err)
	}
}
