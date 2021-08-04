package version

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Implementation for VersionManager
type Ubuntu struct{}

// Get all versions installed on the computer
func (system Ubuntu) GetLocalVersions() []string {

	// Get installed packages
	out, err := exec.Command("apt", "list", "--installed").Output()

	if err != nil {
		return []string{}
	}

	// Filter non php packages
	re := regexp.MustCompile(`(php)([0-9\.]+)(-cli)`)
	packages := re.FindAllString(string(out), -1)

	// Get only version numbers
	var versions []string
	re2 := regexp.MustCompile(`([^0-9\.]+)`)
	for i := 0; i < len(packages); i++ {
		versions = append(versions, re2.ReplaceAllString(packages[i], ""))
	}

	return versions
}

// Get all versions available on the repsystemitory
func (system Ubuntu) GetRemoteVersions() []string {

	// Get remote available packages
	out, err := exec.Command("apt-cache", "search", "^(php)([0-9.]+)$").Output()

	if err != nil {
		return []string{}
	}

	// Transform the output
	re2 := regexp.MustCompile(`([\s]*-[^\n]+)`)
	packages := strings.Fields(re2.ReplaceAllString(string(out), ""))

	// Get only version numbers
	var versions []string
	re3 := regexp.MustCompile(`([^0-9\.]+)`)
	for i := 0; i < len(packages); i++ {
		versions = append(versions, re3.ReplaceAllString(packages[i], ""))
	}

	return versions
}

// Install a version from remote
func (system Ubuntu) InstallVersion(v string) (bool, string) {
	// Check availability
	var versions []string = system.GetRemoteVersions()
	var found bool = false

	for i := range versions {
		if v == versions[i] {
			found = true
		}
	}

	if !found {
		return false, "Required version not found"
	}

	// Install the version
	_, err := exec.Command("apt-get", "install", "--yes", "php"+v).Output()
	if err != nil {
		return false, "Failed installing the version"
	}

	return true, "Version installed successfully"
}

// Remove all packages of a local version
func (system Ubuntu) UninstallVersion(v string) (bool, string) {
	// Check installed version
	var versions []string = system.GetLocalVersions()
	var found bool = false

	for i := range versions {
		if v == versions[i] {
			found = true
		}
	}

	if !found {
		return false, "Local version not found"
	}

	// Remove the version
	_, err := exec.Command("apt-get", "remove", "--purge", "--yes", "php"+v+"-*").Output()
	if err != nil {
		return false, "Failed removing the version"
	}

	return true, "Version removed successfully"
}

// Switch the version to use
func (system Ubuntu) UseVersion(v string) (bool, string) {
	// Check availability
	var versions []string = system.GetLocalVersions()
	var found bool = false

	for i := range versions {
		if v == versions[i] {
			found = true
		}
	}

	if !found {
		return false, "Local version not found"
	}

	// Set the symlink to the version
	var binPath string = "/usr/bin"
	errRemove := os.Remove(binPath + "/php.default")
	if errRemove != nil {
		return false, errRemove.Error()
	}

	errSymlink := os.Symlink(binPath+"/php"+v, binPath+"/php.default")
	if errSymlink != nil {
		return false, errSymlink.Error()
	}

	errChmod := os.Chmod(binPath+"/php.default", 0755)
	if errChmod != nil {
		return false, errChmod.Error()
	}

	return true, "Version switched successfully to " + v
}
