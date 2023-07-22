package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func CallDonate(subscriptionLink string, setting Setting) {
	fmt.Println("curl Donate...")

	var address string
	var donateURL string

	if setting.DonateURL == "stop" {
		address = "https://yebekhe.000webhostapp.com/donate/"
		donateURL = address + "?remove=true" + "&username=" + setting.ChannelName
	} else if setting.DonateURL == "yebekhe" {
		address = "https://yebekhe.000webhostapp.com/donate/"
		donateURL = address + "?url=" + url.QueryEscape(subscriptionLink) + "&username=" + setting.ChannelName
	}

	fmt.Println(donateURL)

	err := curlFunc(donateURL)
	if err != nil {
		fmt.Println(err)
	}

}

func curlFunc(address string) error {

	// Encode the URL
	encodedURL, err := url.Parse(address)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}

	// Make the GET request
	resp, err := http.Get(encodedURL.String())
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer resp.Body.Close()

	return nil

}
