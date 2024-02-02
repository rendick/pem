package main

import (
	"encoding/json"
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

func PackageName() []string {
	ReadJson, err := os.Open("packages.json")
	if err != nil {
		log.Fatal(err)
	}
	defer ReadJson.Close()

	ByteValue, _ := ioutil.ReadAll(ReadJson)
	var packages Packages
	err = json.Unmarshal(ByteValue, &packages)
	if err != nil {
		log.Fatal(err)
	}
	var packageNames []string
	for i := 0; i < len(packages.Packages); i++ {
		packageNames = append(packageNames, packages.Packages[i].Name)
	}

	return packageNames
}
