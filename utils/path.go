package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func AddToPath(newPath string) error {
	if pathExistsInEnv("PATH", newPath) {
		fmt.Printf("The directory %s is already in the PATH.\n", newPath)
		return nil
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("error opening registry key: %v", err)
	}
	defer key.Close()

	currentPath, _, err := key.GetStringValue("PATH")
	if err != nil {
		return fmt.Errorf("error reading PATH value: %v", err)
	}

	newPath = fmt.Sprintf("%s;%s", newPath, currentPath)
	err = key.SetStringValue("PATH", newPath)
	if err != nil {
		return fmt.Errorf("error updating PATH value: %v", err)
	}

	return nil
}

func pathExistsInEnv(envVar, pathToAdd string) bool {
	paths := os.Getenv(envVar)
	for _, p := range filepath.SplitList(paths) {
		if strings.EqualFold(p, pathToAdd) {
			return true
		}
	}
	return false
}
