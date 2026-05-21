package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func AddToPath(newPath string, userInstall bool) error {
	root := registry.LOCAL_MACHINE
	keyPath := `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
	scope := "system"
	if userInstall {
		root = registry.CURRENT_USER
		keyPath = `Environment`
		scope = "user"
	}

	key, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("error opening %s environment registry key: %v", scope, err)
	}
	defer key.Close()

	currentPath, valueType, err := key.GetStringValue("Path")
	if err != nil && err != registry.ErrNotExist {
		return fmt.Errorf("error reading %s Path value: %v", scope, err)
	}
	if err == registry.ErrNotExist {
		valueType = registry.SZ
	}

	if pathExistsInValue(currentPath, newPath) {
		fmt.Printf("The directory %s is already in the %s PATH.\n", newPath, scope)
		return nil
	}

	if currentPath != "" {
		newPath = fmt.Sprintf("%s;%s", newPath, currentPath)
	}
	if err := setPathValue(key, newPath, valueType); err != nil {
		return fmt.Errorf("error updating %s Path value: %v", scope, err)
	}

	return NotifyEnvironmentChanged()
}

func SetSystemEnv(name, value string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("error opening system environment registry key: %v", err)
	}
	defer key.Close()

	if err := key.SetStringValue(name, value); err != nil {
		return fmt.Errorf("error setting system environment variable %s: %v", name, err)
	}
	return NotifyEnvironmentChanged()
}

func setPathValue(key registry.Key, value string, valueType uint32) error {
	if valueType == registry.EXPAND_SZ {
		return key.SetExpandStringValue("Path", value)
	}
	return key.SetStringValue("Path", value)
}

func NotifyEnvironmentChanged() error {
	user32 := windows.NewLazySystemDLL("user32.dll")
	sendMessageTimeout := user32.NewProc("SendMessageTimeoutW")
	environment := windows.StringToUTF16Ptr("Environment")
	var result uintptr

	ret, _, err := sendMessageTimeout.Call(
		0xffff,
		0x001a,
		0,
		uintptr(unsafe.Pointer(environment)),
		0x0002,
		5000,
		uintptr(unsafe.Pointer(&result)),
	)
	if ret == 0 {
		return fmt.Errorf("error notifying environment change: %v", err)
	}
	return nil
}

func pathExistsInValue(paths, pathToAdd string) bool {
	for _, p := range filepath.SplitList(paths) {
		if strings.EqualFold(p, pathToAdd) {
			return true
		}
	}
	return false
}
