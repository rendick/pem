package cmd

import (
	"os"
)

// sudo chmod 777 /var/log/pem

var (
	// path
	logpath      = "/var/log/pem/"
	packagespath = "./packages.json"
)

func Logs() {
	_, err := os.Stat(logpath)
	if os.IsNotExist(err) {
		err := os.Mkdir(logpath, 0750)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(logpath+"pem.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, writeErr := f.WriteString(Log)
	if writeErr != nil {
		panic(writeErr)
	}
}
