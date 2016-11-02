package main

import (
	"log"
	"math/rand"
	"os"
)

// randomStr generate random string for the google Oauth part.
func randomStr() string {
	var bytes = make([]byte, 10)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphaNum[b%byte(len(alphaNum))]
	}
	return string(bytes)
}

// sanityCheck check all mandatory variables.
func sanityCheck() {
	langoustePort := os.Getenv("langoustePort")
	ClientID := os.Getenv("ClientID")
	ClientSecret := os.Getenv("ClientSecret")

	if langoustePort == "" {
		log.Fatal("You must specify a langoustePort.")
	}

	if ClientID == "" {
		log.Fatal("You must specify a google ClientID.")
	}

	if ClientSecret == "" {
		log.Fatal("You must specify a google ClientSecret.")
	}
}
