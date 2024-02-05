package config

import "fmt"

var (
	// style
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Help() {
	fmt.Println(Bold + "Usage: pem [OPTION] [FILE]" + Reset +
		"\n\ninstall: installation of the entered package" +
		"\nlist: printing of all available packages" +
		"\n\npem: <https://github.com/rendick/pem/>")
}

func Missing() {
	fmt.Println("pem: missing arguments\nUsage: pem [OPTION]\nTry --help for more information")
}
