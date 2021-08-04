package global

import (
	"strings"
)

type menuEntry struct {
	title       string
	description string
}

// Get the help menu
func (g Global) GetHelp() string {

	var menu []menuEntry = []menuEntry{
		{"init", "Update environment to use phpenv correctly"},
		{"list", "List all installed versions"},
		{"list-remote", "List all installable versions"},
		{"install", "Install a specific version of PHP"},
		{"uninstall", "Uninstall a specific version of PHP"},
		{"use", "Switch version to use"},
		{"help", "Show this help"},
	}
	var result string

	// Join parts
	result += "Usage: phpenv <command> [<options>]\n\n"
	result += "Commands:\n"
	for i := 0; i < len(menu); i++ {
		result += menu[i].title + strings.Repeat(" ", 20-len(menu[i].title)) + menu[i].description + "\n"
	}

	return result
}
