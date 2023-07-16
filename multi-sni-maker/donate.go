package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func CallDonate(subscriptionLink string, setting Setting) {
	fmt.Println("curl Donate...")

	// make GET request to API to get user by ID
	donateURL := setting.DonateURL + "?url=" + url.QueryEscape(subscriptionLink) + "&username=" + setting.ChannelName

	fmt.Println(donateURL)

	// Encode the URL
	encodedURL, err := url.Parse(donateURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}

	// Make the GET request
	resp, err := http.Get(encodedURL.String())
	if err != nil {
		fmt.Println("Error making GET request:", err)
	}
	defer resp.Body.Close()

}
