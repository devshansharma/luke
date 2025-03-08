package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/devshansharma/luke/internal/models"
	"github.com/devshansharma/luke/pkg/utils"
)

type CollectionDetailsConfig struct {
	Name       string `validate:"required"`
	FolderName string `validate:"omitempty"`
}

// CollectionDetails to list collection items and folders
func CollectionDetails(cfg *CollectionDetailsConfig) error {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return err
	}

	collectionFileName := fmt.Sprintf("%s/%s", dir, getCollectionFileName(cfg.Name))
	_, err = os.Stat(collectionFileName)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("collection does not exist")
	}

	data, err := os.ReadFile(collectionFileName)
	if err != nil {
		return fmt.Errorf("failed to get collection: %s", err.Error())
	}

	obj := models.Collection{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return fmt.Errorf("failed to parse collection: %s", err.Error())
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

	if cfg.FolderName == "" {
		for _, item := range obj.Items {
			fmt.Fprintf(os.Stdout, "- %s\n", item.Name)
		}
	}

	return nil
}
