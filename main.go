package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	config "rendick/pem/settings"
	"strings"
)

// func ArgumentName() {
//
// }

func main() {
	config.Config()
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "install":
			fmt.Printf("%s: missing package\nUsage: pem install [PACKAGE]\nTry --help for more information\n", os.Args[1])
		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "install":
			fmt.Println("intall 2")
			PackageName()
			var packages Packages
			for i := 0; i < len(packages.Packages); i++ {
				fmt.Printf("%s\n", packages.Packages[i].Name)
			}
			if strings.Contains(os.Args[2], "micro") == true {
				micro_cmd, err := exec.Command("sh", "-c", "curl -OL https://github.com/zyedidia/micro/releases/download/v2.0.13/micro-2.0.13-linux64.tar.gz && tar -xvzf micro-2.0.13-linux64.tar.gz && cd micro-2.0.13 && sudo mv micro /usr/local/bin").Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(micro_cmd))
			}

		}
	} else {
		fmt.Printf("pem: missing arguments\nUsage: pem [OPTION]\nTry --help for more information\n")
	}

}
