package utils

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func AddStartupEntry(name, executablePath string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %v", err)
	}
	defer key.Close()

	existingValues, err := key.ReadValueNames(-1)
	if err != nil {
		return fmt.Errorf("failed to read registry values: %v", err)
	}

	for _, v := range existingValues {
		if v == name {
			return nil // Entry already exists, so return without adding again
		}
	}
	err = key.SetStringValue(name, executablePath)
	if err != nil {
		return fmt.Errorf("failed to write to registry: %v", err)
	}

	return nil
}
