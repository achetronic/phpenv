package repository

import "os/exec"

// Implementation for RepositoryManager
type Ubuntu struct{}

// Add PPA repo to the system
func (system Ubuntu) AddRepository() (bool, string) {

	//
	updated, msg := system.UpdatePackages()
	if !updated {
		return false, msg
	}

	// Add the PPA dependencies
	_, errDependencies := exec.Command("apt-get", "install", "--yes", "software-properties-common").Output()
	if errDependencies != nil {
		return false, "Failed installing PPA dependencies"
	}

	// Add PPA repository
	_, errPPA := exec.Command(
		"add-apt-repository", "--yes", "ppa:ondrej/php").Output()
	if errPPA != nil {
		return false, "Failed adding PPA repository"
	}

	return true, "Repository added successfully"
}

// Update packages list
func (system Ubuntu) UpdatePackages() (bool, string) {
	_, err := exec.Command("apt-get", "update").Output()
	if err != nil {
		return false, "Failed updating package list"
	}
	return true, "Package list is up-to-date"
}
