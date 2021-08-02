package hosts

import "runtime"

const UnixHostsPath = "C:/Windows/System32/drivers/etc/hosts"
const WindowsHostsPath = "C:/Windows/System32/drivers/etc/hosts"

func getHostsPath() string {
	if runtime.GOOS == "windows" {
		return WindowsHostsPath
	} else {
		return UnixHostsPath
	}
}
