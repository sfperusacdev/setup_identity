package utils

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func AddStartupEntry(name, executablePath string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %v", err)
	}
	defer key.Close()
	// Establecer el valor del registro para la entrada de inicio automático
	err = key.SetStringValue(name, executablePath)
	if err != nil {
		return fmt.Errorf("failed to set registry value: %v", err)
	}
	return nil
}
