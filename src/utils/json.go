package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func ParseFile(path string, dataType interface{}) {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read config file")
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonFile, dataType)

	if err != nil {
		fmt.Println("Failed to parse config file")
		log.Fatal(err)
	}
}
