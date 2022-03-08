package cmd

import (
	"log"
)

func exitOnError(err error) bool {
	if err != nil {
		log.Fatal(err)
	}
	return true
}