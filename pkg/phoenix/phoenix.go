package phoenix

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Migrate(projectPath, newLocation string) error {
	// Validate input URL.
	newLocaltionURL, err := url.Parse(newLocation)
	if err != nil {
		return fmt.Errorf("can't parse the given new location : %w", err)
	}
	newModuleName := newLocaltionURL.String()

	// Read old module name.
	goModPath := filepath.Join(projectPath, "go.mod")
	oldModuleName, err := readGoModuleName(goModPath)
	if err != nil {
		return fmt.Errorf("can't read go module name : %w", err)
	}

	// Use the new module name in go.mod file.
	if err := ReplaceStringIn(goModPath, "go.mod.tmp", oldModuleName, newModuleName); err != nil {
		return fmt.Errorf("can't replace the module name : %w", err)
	}

	// Replace the module name in all go files in the projectPath.
	if err := ReplaceStringInFiles(projectPath, oldModuleName, newLocation); err != nil {
		return fmt.Errorf("can't iterate over all go file of the projectPath : %w", err)
	}

	return nil
}

func readGoModuleName(goModPath string) (string, error) {
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return "", fmt.Errorf("No go.mod found in given directory. Error : %w", err)
	}

	goMod, err := os.Open(goModPath)
	if err != nil {
		return "", fmt.Errorf("can't open project go mod : %w", err)
	}
	defer goMod.Close()

	scanner := bufio.NewScanner(goMod)
	scanner.Scan()
	goModFirstLine := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("can't read project go mod first line : %w", err)
	}

	return strings.Split(goModFirstLine, " ")[1], nil
}
