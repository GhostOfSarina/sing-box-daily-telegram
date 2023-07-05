package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func SaveSubscribe(filename string, StringConfigAll string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(StringConfigAll)
	if err2 != nil {
		log.Fatal(err2)
	}

}

func ReadFile() (currentReality RealityJson) {

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

func botToken() string {
	dat, err := os.ReadFile("./bot_token.txt")
	if err != nil {
		log.Fatal("error during the ReadFile")
	}
	botToken := string(dat)

	botToken = strings.TrimSpace(botToken)

	return botToken
}

func chatId() string {
	dat, err := os.ReadFile("./chat_id.txt")
	if err != nil {
		log.Fatal("error during the ReadFile")
	}
	chatId := string(dat)

	chatId = strings.TrimSpace(chatId)

	return chatId
}
