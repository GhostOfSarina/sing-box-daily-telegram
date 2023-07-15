package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

func DoSubscribe(setting Setting, StringConfigAll string) (subscriptionNameLink string, err error) {

	if setting.DynamicSubscription {
		randomUUID := rand.Intn(10000)
		subscriptionNameLink = "subscribe." + strconv.Itoa(randomUUID) + ".txt"
	} else {
		subscriptionNameLink = "subscribe.txt"
	}

	SaveSubscribe("./"+subscriptionNameLink, StringConfigAll)

	_, err = exec.Command("/bin/sh", "./make-subscribe.sh").Output()
	if err != nil {
		fmt.Printf("error make-subscribe %s", err)
		return "", err
	}

	return subscriptionNameLink, err

}

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
