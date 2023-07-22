package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("restarting...")
	//Reinstall sing box
	_, err := exec.Command("/bin/sh", "./reinstall-sing-box.sh").Output()
	if err != nil {
		fmt.Printf("error make-subscribe %s", err)
	}

	fmt.Println("read setting file...")

	//Read files
	serverIP := GetOutboundIP().String()

	//read configuration from setting file
	setting, err := ReadSettingsFile()
	if err != nil {
		fmt.Printf("error read settings file %v ", err)
	}

	//todo read setting from config
	var newReality RealityJson

	fmt.Println("renew configuration...")

	//renew existing reality json
	StringConfigZero, StringConfigAll, newReality := RenewConfigurations(setting, serverIP, newReality)

	//block bad websites
	newReality = Block(newReality)

	//save new Reality in file
	err = WriteFile("./reality.json", newReality)
	if err != nil {
		log.Fatal("error during the WriteFile")
	}

	var AggregateSubscriptionNameLink string
	//aggregate subscriptions
	if len(setting.AggregateSubscriptions) > 0 {
		AggregateSubscriptionNameLink = AggregateSubscriptions(setting, StringConfigAll)
	}

	//make subscription
	subscriptionNameLink, err := DoSubscribe(setting, StringConfigAll)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	//call Telegram
	if len(setting.BotToken) > 4 && len(setting.ChatID) > 4 {
		err = CallTelegram(StringConfigZero, setting)
		if err != nil {
			fmt.Printf("error %s", err)
		}

		if len(setting.AggregateSubscriptions) > 0 {
			err = CallTelegram("Aggregate link is: \nhttp://"+serverIP+"/"+AggregateSubscriptionNameLink, setting)
			if err != nil {
				fmt.Printf("error %s", err)
			}
		} else {
			err = CallTelegram("You can also use this link to subscribe to all configuration:\nhttp://"+serverIP+"/"+subscriptionNameLink, setting)
			if err != nil {
				fmt.Printf("error %s", err)
			}
		}
	}

	//call donate endpoint
	if len(setting.DonateURL) > 4 {

		if len(setting.AggregateSubscriptions) > 0 {

			subscriptionNameLinkFull := "http://" + serverIP + "/" + AggregateSubscriptionNameLink
			CallDonate(subscriptionNameLinkFull, setting)

		} else {

			subscriptionNameLinkFull := "http://" + serverIP + "/" + subscriptionNameLink
			CallDonate(subscriptionNameLinkFull, setting)
		}

	}

	//send vnstat file
	if setting.SendVNstat {

		logByte, err := os.ReadFile("./log.txt")
		if err != nil {
			fmt.Printf("error %s", err)
		}
		stringLogByte := string(logByte)
		err = CallTelegram(stringLogByte, setting)
		if err != nil {
			fmt.Printf("error %s", err)
		}
	}

}
