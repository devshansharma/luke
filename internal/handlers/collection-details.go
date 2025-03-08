package handlers

import (
	"fmt"
	"os"
	"strings"
)

type CollectionDetailsConfig struct {
	Name       string `validate:"required"`
	FolderName string `validate:"omitempty"`
}

// CollectionDetails to list collection items and folders
func CollectionDetails(cfg *CollectionDetailsConfig) error {
	obj, _, err := getCollection(cfg.Name)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Collection: %s\n", obj.Info.Name)

	for _, folder := range obj.Folders {
		fmt.Fprintf(os.Stdout, "[] %s\n", folder.Name)

		if strings.EqualFold(cfg.FolderName, folder.Name) {
			for _, item := range folder.Items {
				fmt.Fprintf(os.Stdout, "   - %s\n", item.Name)
			}
		}
	}

	for _, item := range obj.Items {
		fmt.Fprintf(os.Stdout, "- %s\n", item.Name)
	}

	return nil
}
