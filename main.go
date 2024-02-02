package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// func ArgumentName() {
//
// }

var Log string

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "install":
			Log = fmt.Sprintf("(%s)\n%s: missing package\nUsage: pem install [PACKAGE]\nTry --help for more information\n", time.Now().Format("2017-09-07 17:06:06"), os.Args[1])
			fmt.Println("Log content: ", Log)
			Logs()
			os.Exit(0)
			// other commands
		case "help":
			fmt.Println(help)
			os.Exit(0)

		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		// main commands
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
