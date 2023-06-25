package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {

	currentReality := ReadFile()
	currentInbound := currentReality.Inbounds[0]

	var newReality RealityJson
	newReality.Log = currentReality.Log
	newReality.Outbounds = currentReality.Outbounds

	publicKey := getPublicKey()

	serverIP := GetOutboundIP().String()

	StringConfigAll := ""

	ports := []int{443, 8081, 8082, 8083, 8084, 8085, 8086, 8087, 8088, 8089}
	domains := []string{"www.datadoghq.com",
		"apppash.ir",
		"rr2.ir",
		"bazarmag.ir",
		"shiralinia.ir",
		"000webhost.ir",
		"favakar.ir",
		"veket.ir",
		"bing.com",
		"tarhpro.ir"}

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

	//save new Reality in file
	err := WriteFile("./reality.json", newReality)
	if err != nil {
		log.Fatal("error during the WriteFile")
	}

	SaveSubscribe("./subscribe.txt", StringConfigAll)

	cmd, err := exec.Command("/bin/sh", "./multi-sni-maker/make-subscribe.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	fmt.Println(output)

	CallTelegram(StringConfigZero)

}

func CallTelegram(severLink string) error {
	fmt.Println("curl Telegram...")

	botToken := botToken()
	chatId := chatId()

	// make GET request to API to get user by ID
	telegramUrl := "https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=" + chatId + "&text=" + severLink

	request, err := http.NewRequest("GET", telegramUrl, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(err)
		return err
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(responseBody)

	// clean up memory after execution
	defer response.Body.Close()

	return nil
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

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
