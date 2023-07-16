package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadSettingsFile() (setting Setting, err error) {

	// Open our jsonFile
	jsonFile, err := os.Open("./setting.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return setting, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return setting, err

	}

	// we unmarshal our byteArray which contains our
	err = json.Unmarshal(byteValue, &setting)
	if err != nil {
		fmt.Println(err)
		return setting, err

	}

	if len(setting.Ports) != len(setting.Domains) { //|| len(setting.GRPC) != len(setting.Ports)
		log.Fatal("length ports and domain is not equals")
		return setting, fmt.Errorf("length ports and domain")
	}

	return setting, nil

}
