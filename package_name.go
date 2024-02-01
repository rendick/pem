package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Packages struct {
	Packages []Package `json:"package"`
}

type Package struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func PackageName() {
	ReadJson, err := os.Open("./packages.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened packages.json")
	defer ReadJson.Close()

	ByteValue, _ := ioutil.ReadAll(ReadJson)
	var packages Packages
	err = json.Unmarshal(ByteValue, &packages)
	if err != nil {
		log.Fatal(err)
	}

}
