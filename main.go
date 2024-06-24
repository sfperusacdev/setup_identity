package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"setup/utils"

	"github.com/kardianos/osext"
)

func printError(err string) {
	fmt.Println(err)
	fmt.Println("Press Enter ...")
	bufio.NewScanner(os.Stdin).Scan()
	os.Exit(1)
}

func main() {
	basepath, err := osext.ExecutableFolder()
	if err != nil {
		printError(err.Error())
	}
	fmt.Println("base path:", basepath)
	if err := utils.AddToPath(path.Join(basepath, "tools")); err != nil {
		printError(err.Error())
	}
	fmt.Println("tools added to PATH ✅")
	var exepath = path.Join(basepath, "sf_updates_manager.exe")
	if err := utils.AddStartupEntry("SF Update Manager", exepath); err != nil {
		printError(err.Error())
	}
	fmt.Println("sf_updates_manager.exe added to Start ✅")
	if err := installMkcert(); err != nil {
		printError(err.Error())
	}
	fmt.Println("mkcert certificates installed successfully ✅")
	fmt.Println("Installation completed 👍")
	fmt.Println("Press Enter ...")
	bufio.NewScanner(os.Stdin).Scan()
}

func installMkcert() error {
	cmd := exec.Command("mkcert", "-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running mkcert -install: %v", err)
	}
	return nil
}
