package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/devshansharma/luke/internal/models"
	"github.com/devshansharma/luke/pkg/utils"
	"github.com/google/uuid"
)

type AddItemConfig struct {
	CollectionName string `validate:"required,min=3,max=30"`
	ItemName       string `validate:"required,min=3,max=90"`
	FolderName     string `validate:"omitempty,min=3,max=30"`
}

// To add folder or item
func AddItem(cfg *AddItemConfig) error {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return err
	}

	collectionFileName := fmt.Sprintf("%s/%s", dir, getCollectionFileName(cfg.CollectionName))
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

	item := models.Item{
		ID:      uuid.NewString(),
		Name:    cfg.ItemName,
		IsValid: false,
	}

	// add item to collection
	if cfg.FolderName == "" {
		obj.Items = append(obj.Items, item)

		return writeToFile(obj, collectionFileName)
	}

	// add to existing folder, or create one
	for i, f := range obj.Folders {
		if strings.EqualFold(f.Name, cfg.FolderName) {
			obj.Folders[i].Items = append(obj.Folders[i].Items, item)
		}
	}

	return writeToFile(obj, collectionFileName)
}

func writeToFile(obj models.Collection, fileName string) error {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	fileData := string(data) + "\n"
	err = os.WriteFile(fileName, []byte(fileData), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}
