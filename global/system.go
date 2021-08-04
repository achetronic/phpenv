package global

import (
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

// Implementation for VersionManager
type Global struct{}

// Get the short name of the os
func (g Global) GetOS() string {

	switch os := runtime.GOOS; os {

	case "darwin":
		return os
	case "windows":
		return os
	case "linux":
		out, err := exec.Command("cat", "/etc/os-release").Output()
		if err != nil {
			return "linux"
		}

		re := regexp.MustCompile(`(ID=){1}([a-zA-Z_0-9]+)`)
		system := re.FindString(string(out))

		if system == "" {
			return "linux"
		}

		system = strings.TrimPrefix(system, "ID=")
		return strings.ToLower(system)
	default:
		return "other"
	}
}
