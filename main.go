package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"path/filepath"
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
	home, err := os.UserHomeDir()
	if err != nil {
		printError(err.Error())
	}
	basepath, err := osext.ExecutableFolder()
	if err != nil {
		printError(err.Error())
	}
	fmt.Println("home:", home)
	fmt.Println("base path:", basepath)
	if err := utils.AddToPath(filepath.Join(basepath, "tools")); err != nil {
		printError(err.Error())
	}
	fmt.Println("\t-tools added to PATH ✅")
	var exepath = filepath.Join(basepath, "sf_updates_manager.exe")
	if err := utils.AddStartupEntry("SF Update Manager", exepath); err != nil {
		printError(err.Error())
	}
	fmt.Println("\t-sf_updates_manager.exe added to Start ✅")
	if err := installMkcert(basepath); err != nil {
		printError(err.Error())
	}
	fmt.Println("\t-mkcert certificates installed successfully ✅")
	fmt.Println("Installation completed 👍")
	fmt.Println("Press Enter ...")
	bufio.NewScanner(os.Stdin).Scan()
}

func installMkcert(basePath string) error {
	cmd := exec.Command(filepath.Join(basePath, "tools/mkcert"), "-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running mkcert -install: %v", err)
	}
	return nil
}
