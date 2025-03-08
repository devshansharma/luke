package handlers

import (
	"strings"

	"github.com/devshansharma/luke/internal/models"
	"github.com/google/uuid"
)

type AddItemConfig struct {
	Name       string `validate:"required,min=3,max=30"`
	ItemName   string `validate:"required,min=3,max=90"`
	FolderName string `validate:"omitempty,min=3,max=30"`
}

// To add folder or item
func AddItem(cfg *AddItemConfig) error {
	obj, fileName, err := getCollection(cfg.Name)
	if err != nil {
		return err
	}

	item := models.Item{
		ID:      uuid.NewString(),
		Name:    cfg.ItemName,
		IsValid: false,
	}

	// add item to collection
	if cfg.FolderName == "" {
		obj.Items = append(obj.Items, item)

		return writeToFile(obj, fileName)
	}

	// add to existing folder, or create one
	for i, f := range obj.Folders {
		if strings.EqualFold(f.Name, cfg.FolderName) {
			obj.Folders[i].Items = append(obj.Folders[i].Items, item)
		}
	}

	return writeToFile(obj, fileName)
}
