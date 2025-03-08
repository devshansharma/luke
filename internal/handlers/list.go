package handlers

import (
	"fmt"
	"os"

	"github.com/devshansharma/luke/pkg/utils"
)

// ListCollection for showing list of collections
func ListCollection() error {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return err
	}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if len(filterCollection(dirEntries)) > 0 {
		fmt.Fprintln(os.Stdout, "Collections found:")
	}

	if len(filterCollection(dirEntries)) == 0 {
		fmt.Fprintln(os.Stderr, "No Collections found")
	}

	for _, entry := range filterCollection(dirEntries) {
		fmt.Fprintf(os.Stdout, "- %s\n", getCollectionName(entry))
	}

	return nil
}

// ListEnvironment list of environments
func ListEnvironment() error {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return err
	}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if len(filterEnvironment(dirEntries)) > 0 {
		fmt.Fprintln(os.Stdout, "Environments found:")
	}

	if len(filterEnvironment(dirEntries)) == 0 {
		fmt.Fprintln(os.Stderr, "No Environments found")
	}

	for _, entry := range filterEnvironment(dirEntries) {
		fmt.Fprintf(os.Stdout, "- %s\n", getEnvName(entry))
	}

	return nil
}
