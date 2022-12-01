package main

import (
	"fmt"

	"github.com/bitfield/script"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	query := fmt.Sprintln(".items[].spec.rules[].host")
	ingresses, err := script.Exec("kubectl get ing -A -ojson").JQ(query).Freq().Column(2).Stdout()
	if err != nil {
		log.Fatal(err)
	}
	if ingresses == 0 {
		log.Info("No resources found")
	}
}
