package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/devshansharma/luke/internal/models"
	"github.com/devshansharma/luke/pkg/utils"
	"github.com/google/uuid"
)

type AddCollectionConfig struct {
	Name string `validate:"required,min=3,max=30"`
}

// To add folder or item
func AddCollection(cfg *AddCollectionConfig) error {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return err
	}

	collectionFile := fmt.Sprintf("%s/%s", dir, getCollectionFileName(cfg.Name))
	_, err = os.Stat(collectionFile)
	if err == nil {
		return fmt.Errorf("collection already exists")
	}

	file, err := os.Create(collectionFile)
	if err != nil {
		return err
	}
	defer file.Close()

	obj := models.Collection{
		Info: models.Info{
			Name: cfg.Name,
			ID:   uuid.NewString(),
		},
	}

	data, err := json.MarshalIndent(obj, "", "   ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(file, string(data))
	if err != nil {
		return err
	}

	return nil
}
