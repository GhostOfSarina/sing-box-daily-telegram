package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func ReadRealityFile() (currentReality RealityJson) {

	// Let's first read the `reality.json` file
	content, err := os.ReadFile("./reality.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshal the data into `currentReality`
	err = json.Unmarshal(content, &currentReality)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return currentReality
}

func WriteFile(filename string, newReality RealityJson) error {

	file, err := json.MarshalIndent(newReality, "", " ")
	if err != nil {
		log.Fatal("Error during MarshalIndent(): ", err)
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Error during WriteFile(): ", err)
		return err
	}

	return nil
}

func getPublicKey() string {
	dat, err := os.ReadFile("./public_key.txt")
	if err != nil {
		log.Fatal("error during the ReadFile")
	}
	publicKey := string(dat)

	publicKey = strings.TrimSpace(publicKey)

	return publicKey
}
