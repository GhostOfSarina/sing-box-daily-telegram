package main

import (
	"strconv"

	"github.com/google/uuid"
)

func RenewConfigurations(setting Setting, serverIP string, newReality RealityJson) (StringConfigZero string, StringConfigAll string, outReality RealityJson) {

	publicKey := getPublicKey()

	currentReality := ReadRealityFile()
	currentInbound := currentReality.Inbounds[0]

	newReality.Outbounds = currentReality.Outbounds

	newReality.Inbounds = make([]Inbound, len(setting.Domains))

	for counter := 0; counter < len(setting.Domains); counter++ {

		inbound := currentInbound
		inbound.ListenPort = setting.Ports[counter]
		inbound.TLS.ServerName = setting.Domains[counter]
		inbound.TLS.Reality.Handshake.Server = setting.Domains[counter]

		name := setting.ChannelName + "-" + setting.Domains[counter]
		inbound.Users = []User{
			{
				Name: name,
				UUID: uuid.New().String(),
				Flow: "xtls-rprx-vision",
			},
		}

		newReality.Inbounds[counter] = inbound

		//GRPC setting
		// if setting.GRPC[counter] {

		// 	inbound.Users[0].Flow = ""
		// 	inbound.Users[0].Name = ""

		// 	newReality.Inbounds[counter].Transport = new(Transport)
		// 	newReality.Inbounds[counter].Transport.Type = "grpc"
		// 	newReality.Inbounds[counter].Transport.ServiceName = name
		// 	newReality.Inbounds[counter].Transport.IdleTimeout = "15s"
		// 	newReality.Inbounds[counter].Transport.PingTimeout = "15s"
		// 	newReality.Inbounds[counter].Transport.PermitWithoutStream = false
		// }

		//capture setting

		StringConfig := "vless://" + inbound.Users[0].UUID + "@" + serverIP + ":" + strconv.Itoa(setting.Ports[counter]) +
			"?encryption=none&flow=xtls-rprx-vision&security=reality&sni=" + setting.Domains[counter] +
			"&fp=chrome&pbk=" + publicKey + "&sid=" + inbound.TLS.Reality.ShortID[0] + "&type=tcp&headerType=none#" + name

		if counter == 0 {
			StringConfigZero = StringConfig
		}

		StringConfigAll += StringConfig + "\n"

	}

	outReality = newReality
	return StringConfigZero, StringConfigAll, outReality

}
