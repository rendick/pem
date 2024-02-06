package cmd

// alpha version

import (
	"fmt"
	"os"

	config "github.com/rendick/pem/settings"
)

var Remove string

func UninstallPacage(packageName string) {
	Remove = fmt.Sprintf("/bin/%s", packageName)
	uninstall := os.Remove(Remove)
	if uninstall != nil {
		fmt.Printf(config.Red+config.Bold+"%s not successfully uninstalled!\nCould not find file!\n"+config.Reset, packageName)
	} else {
		fmt.Printf(config.Green+config.Bold+"%s successfully uninstalled!\n"+config.Reset, packageName)
	}
}
