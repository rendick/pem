package cmd

import (
	"fmt"
	"log"
	"os"

	config "github.com/rendick/pem/settings"
)

func PackageStats() {
	stats, err := os.ReadDir(config.PackageDir)
	if err != nil {
		log.Fatal(err)
	}

	for num, files := range stats {
		files_type, err := os.Stat(config.PackageDir + files.Name())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(config.Bold+"%d:"+config.Reset+" %s "+config.Bold+"%s |"+config.Reset+config.Bold+" (%d MiB)\n"+config.Reset,
			num,
			files,
			files_type.ModTime(),
			files_type.Size()/1024)
	}
}
