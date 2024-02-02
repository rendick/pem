package main

import (
	"fmt"
	"os"
	"time"
)

var Log string

func Switch() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "install", "-I", "-i", "i": // main commands
			Log = fmt.Sprintf(Red+"(%s)"+Reset+"\n%s: missing package\nUsage: pem install [PACKAGE]\nTry --help for more information\n", time.Now().Format("2006-01-02 15:04:05"), os.Args[1])
			Logs()
			os.Exit(0)

		case "list", "-L", "-l", "l": // main commands
			PackageName()

		case "help", "-H", "-h", "h": // other commands
			fmt.Println(help)
			os.Exit(0)

		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "install": // main commands
			InstallPackage(os.Args[2])

		}
	} else {
		fmt.Printf("pem: missing arguments\nUsage: pem [OPTION]\nTry --help for more information\n")
	}
}
