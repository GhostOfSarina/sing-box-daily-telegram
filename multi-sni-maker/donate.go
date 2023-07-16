package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func CallDonate(severLink string, setting Setting) error {
	fmt.Println("curl Donate...")

	// make GET request to API to get user by ID
	donateURL := setting.DonateURL + "?data=" + url.QueryEscape(severLink)

	// fmt.Println(donateURL)

	// Encode the URL
	encodedURL, err := url.Parse(donateURL)
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
