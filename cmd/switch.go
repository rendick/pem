package cmd

import (
	"fmt"
	"os"

	config "github.com/rendick/pem/settings"
)

var Log string

var (
	InstallLog = "\n%s: missing package\nUsage: pem install [PACKAGE]\nTry --help for more information\n"
	RemoveLog  = "\n%s: missing package\nUsage: pem remove [PACKAGE]\nTry --help for more information\n"
)

func Switch() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "install", "-I", "-i", "i": // main command
			Log = fmt.Sprintf(config.Red+"(%s)"+config.Reset+InstallLog, config.Time, os.Args[1])
			fmt.Println(Log)
			Logs()
			os.Exit(0)

		case "list", "-L", "-l", "l": // main command
			Scrapper()

		case "update", "-U", "-u", "u": // main command
			fmt.Println("Soon.")
			os.Exit(0)

		case "remove", "-R", "-r", "r": // main command
			Log = fmt.Sprintf(config.Red+"(%s)"+config.Reset+RemoveLog, config.Time, os.Args[1])
			fmt.Println(Log)
			Logs()
			os.Exit(0)

		case "help", "--help", "-H", "-h", "h": // other command
			config.Help()
			os.Exit(0)
		case "version", "--version", "-V", "-v", "v": // other command
			config.Version()
			os.Exit(0)

		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "install", "-I", "-i", "i": // main command
			InstallPackage(os.Args[2])

		case "remove", "-R", "-r", "r": // main command
			UninstallPacage(os.Args[2])
		}
	} else {
		config.Missing()
	}
}
