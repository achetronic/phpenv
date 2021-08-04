package main

import (
	"fmt"
	"os"
	"phpenv/global"
	"phpenv/repository"
	"phpenv/version"
)

// Init global functions
var globalFunctions global.Global
var osType string = globalFunctions.GetOS()
var args []string = os.Args[1:]

// Init OS dependant functions
var source repository.RepositoryManager
var php version.VersionManager

// Install repository
func InitCommand() {
	fmt.Println("Adding packages repository...")
	_, msg := source.AddRepository()
	fmt.Println(msg)
}

// List all installed versions
func ListLocalCommand() {
	fmt.Println("Listing installed versions...")
	versions := php.GetLocalVersions()
	for i := 0; i < len(versions); i++ {
		fmt.Println(versions[i])
	}
}

// List all installed versions
func ListRemoteCommand() {
	fmt.Println("Listing available versions...")
	versions := php.GetRemoteVersions()
	for i := 0; i < len(versions); i++ {
		fmt.Println(versions[i])
	}
}

// Install version passed by args
func InstallCommand() {
	if len(args) < 2 {
		fmt.Println("Set a version to install")
		return
	}

	fmt.Println("Installing desired version...")
	_, msg := php.InstallVersion(args[1])
	fmt.Println(msg)
}

// Uninstall version passed by args
func UninstallCommand() {
	if len(args) < 2 {
		fmt.Println("Set a version to uninstall")
		return
	}

	fmt.Println("Uninstalling desired version...")
	_, msg := php.UninstallVersion(args[1])
	fmt.Println(msg)
}

// Use version passed by args
func UseCommand() {
	if len(args) < 2 {
		fmt.Println("Set a version to switch to")
		return
	}

	fmt.Println("Switching current version...")
	_, msg := php.UseVersion(args[1])
	fmt.Println(msg)
}

// Execute main code
func main() {

	// OS selector
	switch osType {
	case "ubuntu":
		source = repository.Ubuntu{}
		php = version.Ubuntu{}
	default:
		source = repository.Ubuntu{}
		php = version.Ubuntu{}
	}

	// Action selector
	cmd := "help"
	if len(args) > 0 {
		cmd = args[0]
	}

	switch cmd {
	case "init":
		InitCommand()
	case "list":
		ListLocalCommand()
	case "list-remote":
		ListRemoteCommand()
	case "install":
		InstallCommand()
	case "uninstall":
		UninstallCommand()
	case "use":
		UseCommand()
	default:
		fmt.Print(globalFunctions.GetHelp())
	}
}
