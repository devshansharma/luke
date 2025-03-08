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
	"github.com/spf13/cobra"
)

type AddFolderConfig struct {
	CollectionName string `validate:"required,min=3,max=30"`
	FolderName     string `validate:"required,min=3,max=30"`
}

// To add folder or item
func AddFolder(cfg *AddFolderConfig) error {
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

	for _, folder := range obj.Folders {
		if strings.EqualFold(folder.Name, cfg.FolderName) {
			return fmt.Errorf("folder already exists: %q", cfg.FolderName)
		}
	}

	obj.Folders = append(obj.Folders, models.Folder{
		ID:    uuid.NewString(),
		Name:  cfg.FolderName,
		Items: make([]models.Item, 0),
	})

	data, err = json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal: %s", err)
	}

	fileData := string(data) + "\n"
	err = os.WriteFile(collectionFileName, []byte(fileData), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}

// AddFolderCompletion for adding a completion functionality
func AddFolderCompletion(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return []cobra.Completion{}, cobra.ShellCompDirectiveNoFileComp
	}

	var suggestions = []string{}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return []cobra.Completion{}, cobra.ShellCompDirectiveNoFileComp
	}

	for _, dir := range filterCollection(entries) {
		if strings.HasPrefix(getCollectionName(dir), toComplete) {
			suggestions = append(suggestions, getCollectionName(dir))
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}
