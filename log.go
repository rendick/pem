package main

import (
	"fmt"
	"os"
)

// sudo chmod 777 /var/log/pem

func Logs() {
	_, err := os.Stat("/var/log/pem")
	if os.IsNotExist(err) {
		err := os.Mkdir("/var/log/pem", 0750)
		if err != nil {
			panic(err)
		}
	}

	writeErr := os.WriteFile("/var/log/pem/pem.log", []byte(Log), 0644)
	if writeErr != nil {
		panic(writeErr)
	}
	fmt.Println("Log written successfully.")
}
