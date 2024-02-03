package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

var URL string

func InstallPackage(packageName string) {

	URL = fmt.Sprintf("https://pem-packages.vercel.app/packages/%s.html", os.Args[2])
	fmt.Println(URL)
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer response.Body.Close()

	tempFile, err := os.CreateTemp("", "micro.html")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, response.Body)
	if err != nil {
		fmt.Println("Error copying content to temporary file:", err)
		return
	}

	command := exec.Command("bash", tempFile.Name())

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing the command:", err)
		return
	}

	fmt.Println(string(output))

}