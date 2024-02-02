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

	f, err := os.OpenFile("/var/log/pem/pem.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, writeErr := f.WriteString(Log)
	if writeErr != nil {
		panic(writeErr)
	} else {
		fmt.Println(Green + "Log written successfully." + Reset)
	}
}
