package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func AggregateSubscriptions(setting Setting, StringConfigAll string) string {

	AllConfigurations := StringConfigAll

	for _, link := range setting.AggregateSubscriptions {

		fmt.Println("Curl URL: ", link)
		result, err := getURL(link)
		if err != nil {
			fmt.Println("Error getting aggregate subscription ", err.Error())
			continue
		}
		AllConfigurations = AllConfigurations + "\n" + result + "\n"
	}

	var subscriptionNameLink string

	if setting.DynamicSubscription {
		randomUUID := rand.Intn(10000)
		subscriptionNameLink = "aggregate." + strconv.Itoa(randomUUID) + ".txt"
	} else {
		subscriptionNameLink = "aggregate.txt"
	}

	SaveSubscribe("./"+subscriptionNameLink, AllConfigurations)

	return subscriptionNameLink

}

func getURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %v", err)
	}

	return string(data), nil
}
