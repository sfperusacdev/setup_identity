package utils

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func AddStartupEntry(name, executablePath string, userInstall bool) error {
	root := registry.LOCAL_MACHINE
	scope := "machine"
	if userInstall {
		root = registry.CURRENT_USER
		scope = "user"
	}

	key, err := registry.OpenKey(root, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open %s startup registry key: %v", scope, err)
	}
	defer key.Close()

	err = key.SetStringValue(name, executablePath)
	if err != nil {
		return fmt.Errorf("failed to set %s startup registry value: %v", scope, err)
	}
	return nil
}
