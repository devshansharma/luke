package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/devshansharma/luke/internal/models"
	"github.com/devshansharma/luke/pkg/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type AddFolderConfig struct {
	Name       string `validate:"required,min=3,max=30"`
	FolderName string `validate:"required,min=3,max=30"`
}

// To add folder or item
func AddFolder(cfg *AddFolderConfig) error {
	obj, fileName, err := getCollection(cfg.Name)
	if err != nil {
		return err
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

	return writeToFile(obj, fileName)
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
