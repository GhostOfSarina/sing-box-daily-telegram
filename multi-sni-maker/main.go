package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/google/uuid"
)

func main() {

	//Reinstall sing box
	_, err := exec.Command("/bin/sh", "./reinstall-sing-box.sh").Output()
	if err != nil {
		fmt.Printf("error make-subscribe %s", err)
	}

	//Read files

	currentReality := ReadFile()
	currentInbound := currentReality.Inbounds[0]

	var newReality RealityJson
	newReality.Outbounds = currentReality.Outbounds

	publicKey := getPublicKey()

	serverIP := GetOutboundIP().String()

	StringConfigAll := ""

	ports := []int{443, 22, 23, 3389, 110, 143, 8086, 8087, 8088, 8089, 8090}
	domains := []string{"www.datadoghq.com",
		"000webhost.ir",
		"speedtest.net",
		"speed.cloudflare.com",
		"fruitfulcode.com",
		"favakar.ir",
		"veket.ir",
		"benecke.com",
		"tarhpro.ir",
		"fernandotrueba.com",
		"mathhub.info",
	}

	if len(ports) != len(domains) {
		log.Fatal("length ports and domain is not equals")

	}

	newReality.Inbounds = make([]Inbound, len(domains))
	StringConfigZero := ""
	for i := 0; i < len(domains); i++ {

		inbound := currentInbound
		inbound.ListenPort = ports[i]
		inbound.TLS.ServerName = domains[i]
		inbound.TLS.Reality.Handshake.Server = domains[i]

		inbound.Users = []User{
			{
				NAME: "SB-" + domains[i],
				UUID: uuid.New().String(),
				Flow: "xtls-rprx-vision",
			},
		}

		newReality.Inbounds[i] = inbound

		StringConfig := "vless://" + inbound.Users[0].UUID + "@" + serverIP + ":" + strconv.Itoa(ports[i]) +
			"?encryption=none&flow=xtls-rprx-vision&security=reality&sni=" + domains[i] +
			"&fp=chrome&pbk=" + publicKey + "&sid=" + inbound.TLS.Reality.ShortID[0] + "&type=tcp&headerType=none#SB-" + domains[i]

		if i == 0 {
			StringConfigZero = StringConfig
		}

		StringConfigAll += StringConfig + "\n"

	}

	//add bad websites
	newReality = Block(newReality)

	//save new Reality in file
	err = WriteFile("./reality.json", newReality)
	if err != nil {
		log.Fatal("error during the WriteFile")
	}

	SaveSubscribe("./subscribe.txt", StringConfigAll)

	_, err = exec.Command("/bin/sh", "./make-subscribe.sh").Output()
	if err != nil {
		fmt.Printf("error make-subscribe %s", err)
	}

	err = CallTelegram(StringConfigZero)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	err = CallTelegram("You can also use this link to subscribe to all configuration:\n" + serverIP + "/subscribe.txt")
	if err != nil {
		fmt.Printf("error %s", err)
	}

}
