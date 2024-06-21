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

func main() {
	basepath, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := utils.AddToPath(filepath.Join(basepath, "tools")); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("/tools added to PATH ✅")

	var exepath = filepath.Join(basepath, "sf_updates_manager.exe")
	if err := utils.AddStartupEntry("SF Update Manager", exepath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("sf_updates_manager.exe added to Start ✅")
	if err := installMkcert(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("mkcert certificates installed successfully ✅")
	fmt.Println("Installation completed 👍")
	fmt.Println("Press Enter ...")
	bufio.NewScanner(os.Stdin).Scan()
}

func installMkcert() error {
	cmd := exec.Command("tools/mkcert", "-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running mkcert -install: %v", err)
	}
	return nil
}
