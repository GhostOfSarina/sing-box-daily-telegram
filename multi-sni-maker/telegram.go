package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func CallTelegram(severLink string) error {
	fmt.Println("curl Telegram...")

	botToken := botToken()
	chatId := chatId()

	// make GET request to API to get user by ID
	telegramUrl := "https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=" + chatId + "&text=" + url.QueryEscape(severLink)

	fmt.Println(telegramUrl)

	// Encode the URL
	encodedURL, err := url.Parse(telegramUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return err
	}

	// Make the GET request
	resp, err := http.Get(encodedURL.String())
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
