package cmd

import (
	"os"
)

// sudo chmod 777 /var/log/pem

var (
	// path
	LogPath = "/var/log/pem/"
)

func Logs() {
	_, err := os.Stat(LogPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(LogPath, 0750)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(LogPath+"pem.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, writeErr := f.WriteString(Log)
	if writeErr != nil {
		panic(writeErr)
	}
}
