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

	for index, entry := range filterCollection(dirEntries) {
		fmt.Fprintf(os.Stdout, "%02d %s\n", index+1, getCollectionName(entry))
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

	for index, entry := range filterEnvironment(dirEntries) {
		fmt.Fprintf(os.Stdout, "%02d %s\n", index+1, getEnvName(entry))
	}

	return nil
}
