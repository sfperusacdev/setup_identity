package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"path/filepath"
	"setup/utils"
	"strings"

	"github.com/kardianos/osext"
)

func printError(err string) {
	fmt.Println(err)
	fmt.Println("Press Enter ...")
	bufio.NewScanner(os.Stdin).Scan()
	os.Exit(1)
}

func main() {
	const defaultStartupExecutable = "sf_manager.exe"

	basepath, err := osext.ExecutableFolder()
	if err != nil {
		printError(err.Error())
	}
	fmt.Println("base path:", basepath)

	home, err := os.UserHomeDir()
	if err != nil {
		printError(err.Error())
	}
	userInstall := pathInside(basepath, home)

	if err := utils.AddToPath(filepath.Join(basepath, "tools"), userInstall); err != nil {
		printError(err.Error())
	}
	fmt.Println("\t-tools added to PATH ✅")
	exepath := defaultStartupExecutable
	if len(os.Args) > 1 {
		exepath = os.Args[1]
	}
	if !filepath.IsAbs(exepath) {
		exepath = filepath.Join(basepath, exepath)
	}
	if err := utils.AddStartupEntry("SF Update Manager", exepath, userInstall); err != nil {
		printError(err.Error())
	}
	fmt.Printf("\t-%s added to Start ✅\n", exepath)
	if !userInstall {
		caRoot := filepath.Join(basepath, ".CA")
		if err := os.MkdirAll(caRoot, 0755); err != nil {
			printError(err.Error())
		}
		if err := os.Setenv("CAROOT", caRoot); err != nil {
			printError(err.Error())
		}
		if err := utils.SetSystemEnv("CAROOT", caRoot); err != nil {
			printError(err.Error())
		}
		fmt.Printf("\t-CAROOT set to %s ✅\n", caRoot)
	}
	if err := installMkcert(basepath); err != nil {
		printError(err.Error())
	}
	fmt.Println("\t-mkcert certificates installed successfully ✅")
	fmt.Println("Installation completed 👍")
	fmt.Println("Press Enter ...")
	bufio.NewScanner(os.Stdin).Scan()
}

func pathInside(path, base string) bool {
	rel, err := filepath.Rel(base, path)
	if err != nil {
		return false
	}
	return rel == "." ||
		(!filepath.IsAbs(rel) && rel != ".." &&
			!strings.HasPrefix(rel, ".."+string(filepath.Separator)))
}

func installMkcert(basePath string) error {
	cmd := exec.Command(filepath.Join(basePath, "tools", "mkcert.exe"), "-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running mkcert -install: %v", err)
	}
	return nil
}
