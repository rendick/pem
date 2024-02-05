package cmd

import (
	"fmt"
	"os"
	"time"

	config "github.com/rendick/pem/settings"
)

var Log string

func Switch() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "install", "-I", "-i", "i": // main commands
			Log = fmt.Sprintf(config.Red+"(%s)"+config.Reset+"\n%s: missing package\nUsage: pem install [PACKAGE]\nTry --help for more information", time.Now().Format("2006-01-02 15:04:05"), os.Args[1])
			fmt.Println(Log)
			Logs()
			os.Exit(0)

		case "list", "-L", "-l", "l": // main commands
			Scrapper()

		case "help", "-H", "-h", "h": // other commands
			config.Help()
			os.Exit(0)
		case "update", "-U", "-u", "u":
			fmt.Println("Update...")
		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "install": // main commands
			InstallPackage(os.Args[2])
		}
	} else {
		config.Missing()
	}
}
