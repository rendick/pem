package cmd

// alpha version

import (
	"fmt"
	"log"
	"os"

	config "github.com/rendick/pem/settings"
)

var Remove string

var (
	PackagePath = "/bin/"
	ErrorMsg    = "%s not successfully uninstalled!\nCould not find file!\n"
	SuccessMsg  = "%s successfully uninstalled!\n"
)

func UninstallPacage(packageName string) {
	Remove = fmt.Sprintf("%s%s", PackagePath, packageName)
	uninstall := os.Remove(Remove)
	if uninstall != nil {
		log.Printf(config.Red+config.Bold+"%s not successfully uninstalled! Could not find file!"+config.Reset, packageName)
	} else {
		log.Printf(config.Green+config.Bold+"%s successfully uninstalled!"+config.Reset, packageName)
	}
}
